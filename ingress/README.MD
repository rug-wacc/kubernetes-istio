Install NGINX Ingress Controller with respect to your environment using guide from https://kubernetes.github.io/ingress-nginx/deploy/

Then deploy apple and banana services.

```sh
kubectl apply -f .\apple.yaml
kubectl apply -f .\banana.yaml
```

Then apply the ingress
```sh
kubectl apply -f .\ingress.yaml
```

Now you should be able to call services on http://localhost/apple and http://localhost/banana