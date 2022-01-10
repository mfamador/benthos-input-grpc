# gRPC benthos input


## Create a custom benthos input that receives messages from a gRPC server


### Run server to read messages and deliver them to the benthos pipeline

```
go run cmd/server/main.go
```


### Run the client to post message

```
go run cmd/client/main.go
```

