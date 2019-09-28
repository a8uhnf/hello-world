package main

import (
	"fmt"
	"time"

	"cloud.google.com/go/storage"
)

func main() {
	url, err := genSignedURL1("image/png", "80d16fca-7264-4931-bae7-5b3173aca9a5")
	if err != nil {
		panic(err)
	}
	fmt.Println(url)
}
func genSignedURL1(ct, key string) (string, error) {
	// Generates a signed URL for use in the PUT request to GCS.
	// Generated URL should be expired after 15 mins.
	url, err := storage.SignedURL(uploadableBucket, key, &storage.SignedURLOptions{
		GoogleAccessID: serviceAccountName,
		Method:         "PUT",
		Expires:        time.Now().Add(15 * time.Minute),
		ContentType:    ct,
		// To avoid management for private key, use SignBytes instead of PrivateKey.
		// In this example, we are using the `iam.serviceAccounts.signBlob` API for signing bytes.
		// If you hope to avoid API call for signing bytes every time,
		// you can use self hosted private key and pass it in Privatekey.
		PrivateKey: cfg.PrivateKey,
		// SignBytes: func(b []byte) ([]byte, error) {
		// 	resp, err := iamService.Projects.ServiceAccounts.SignBlob(
		// 		serviceAccountID,
		// 		&iam.SignBlobRequest{BytesToSign: base64.StdEncoding.EncodeToString(b)},
		// 	).Context(r.Context()).Do()
		// 	if err != nil {
		// 		return nil, err
		// 	}
		// 	return base64.StdEncoding.DecodeString(resp.Signature)
		// },
	})
	if err != nil {

		return "", err
	}
	return url, err
}
