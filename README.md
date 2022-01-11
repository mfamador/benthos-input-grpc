# gRPC custom benthos input


## Create a custom benthos input that receives messages from a gRPC server


### Run server to read messages and deliver them to the benthos pipeline

```
go run cmd/server/main.go -c config/pipeline.yaml
```


### Run the client to post random messages

```
go run cmd/client/main.go
```

