# RUN LOCALLY
```sh
go run testhttpserver.go
```

Now you can try http://localhost:9099/hello, http://localhost:9099/headers, and http://localhost:9099/ip.

# RUN ON Kubernetes
```sh
docker build -t httpserver .
kubectl apply -f .\deploy.yaml 
kubectl apply -f .\service.yaml
```

