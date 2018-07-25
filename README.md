Simple node disk manager client

Build binary

```
go build -o ndm-client
```


Build Docker image

```
sudo docker build -t username/ndm-client:0.1 .
```


Deploy NDM client in kubernetes

```
kubectl apply -f ndm-operator.yaml 
kubectl apply -f ndm-client.yaml 
```

Result
```

kubectl logs -f ndm-client-76c4f864d7-4lvsw
Starting NDM Client
disk-ecf109ad98be73104b0c9c77924e47b3
disk-f40ddd5322af3784be2c1cc2b903535a

```



