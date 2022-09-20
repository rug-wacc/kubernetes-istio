# Setup Kubernetes Dashboard
## Deploy Dashboard
Use the following command for deploying the dashboard.
```sh
kubectl apply -f ./recommended.yaml
```

## Access the dashboard
First you need to proxy it,
```sh
kubectl proxy
```
then access it on http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/#/overview

Now you will need a token to log in. For that, we first have to create a user. Note that this user has admin privilege and is only for demo purposes. 
```sh
kubectl apply -f user.yaml
```

Then, using the following command will generate the access token for you. 
```sh
kubectl -n kubernetes-dashboard create token admin-user
```