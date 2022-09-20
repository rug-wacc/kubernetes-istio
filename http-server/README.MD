# RUN LOCALLY
```sh
go run testhttpserver.go
```

Now you can try http://localhost:9099/hello, http://localhost:9099/headers, and http://localhost:9099/ip.
# RUN ON DOCKER
```sh
docker build -t httpserver .
docker images -a
docker run --name myServer -p 80:9099  httpserver 
```
Now you can try http://localhost/hello, http://localhost/headers, and http://localhost/ip.
# RUN ON Kubernetes
```sh
docker build -t httpserver .
kubectl apply -f .\deploy.yaml 
kubectl apply -f .\service.yaml
```

Now that the deployment and the serivce is ready. You have to forward the port from cluster to outside using the following command.
```sh
kubectl port-forward service/http-service 80:9099
```

Since we forwarded on port 80 now you can access the server without specifiying the port and try the followings:

http://localhost/hello, http://localhost/headers, and http://localhost/ip.
