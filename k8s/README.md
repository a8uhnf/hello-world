# Instruction
### To create secret. Secret needs to be in same namespace as deployment.
`kubectl create secret -n kube-system generic db-user-pass --from-file=./username.txt --from-file=./password.txt`

### To deploy `hello-world` app.

`kubectl apply -f k8s/deploy.yaml`
