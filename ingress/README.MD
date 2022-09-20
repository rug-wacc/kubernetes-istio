## Enabling Ingress
Make sure that you stoped all the port binding from port 80. 
```sh
kubectl apply -f enable-ingress.yaml
```
If you are using other kubernetes setup, you can install NGINX Ingress Controller with respect to your environment using guide from https://kubernetes.github.io/ingress-nginx/deploy/
## Deploy resources
Then deploy apple and banana services.

```sh
kubectl apply -f .\apple.yaml
kubectl apply -f .\banana.yaml
```

Then apply the ingress
```sh
kubectl apply -f .\ingress.yaml
```
## Access
Now you should be able to call services on http://localhost/apple and http://localhost/banana