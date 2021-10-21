# Setup Kubernetes Dashboard

```sh
kubectl apply -f ./recommended.yaml
```

## Access the dashboard
First you need to proxy it,
```sh
kube proxy
```
then access it on http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/#/overview
