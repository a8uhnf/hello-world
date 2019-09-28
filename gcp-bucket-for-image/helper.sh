curl -v -X OPTIONS -H "Host: storage.googleapis.com" -H "Access-Control-Request-Method: PUT"  -H "Origin: https://gcs.somedomain.com:8081" "https://storage.googleapis.com/sticker-driver-image/cors.txt"
gsutil cors get gs://bucket-name
gsutil cors set cors-json-file.json gs://bucket-name
