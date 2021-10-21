# ISTIO DEMO
## instal istio
you can visit https://istio.io/latest/docs/setup/getting-started/ for a complete installation guide, or follow these instructions if you are using windows.

```sh
.\istio-1.11.2\bin\istioctl install --set profile=demo -y
kubectl label namespace default istio-injection=enabled
kubectl apply -f ./istio-1.11.2/samples/addons
kubectl rollout status deployment/kiali -n istio-system
```

Then run the dashboard using
```sh
.\istio-1.11.2\bin\istioctl dashboard kiali
```

It is accessable on http://localhost:20001/kiali/console/overview

## Run demo application
```sh
kubectl apply -f bookinfo.yaml
```

Access the webpage on http://localhost:30001/productpage