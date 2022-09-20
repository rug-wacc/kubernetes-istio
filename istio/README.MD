# ISTIO DEMO
## instal istio
you can visit https://istio.io/latest/docs/setup/getting-started/ for a complete installation guide, or follow these instructions if you are using windows.

```sh
.\istio-1.15.0\bin\istioctl install --set profile=demo -y
kubectl label namespace default istio-injection=enabled
kubectl apply -f ./istio-1.15.0/samples/addons
kubectl rollout status deployment/kiali -n istio-system
```

Then run the dashboard using
```sh
.\istio-1.15.0\bin\istioctl dashboard kiali
```

It is accessable on http://localhost:20001/kiali/console/overview

## Run demo application
```sh
kubectl apply -f bookinfo.yaml
kubectl apply -f bookinfo-gateway.yaml
```

Access the webpage on http://localhost/productpage. Visiting this web page will result in traffic to flow through the serviceses. Therefore, you can explore the Kiali UI to see the graph and more. 