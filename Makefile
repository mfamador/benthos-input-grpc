run-server:
	go run cmd/server/main.go -c config/pipeline.yaml

run-server-error:
	go run cmd/server/main.go -c config/pipeline-error.yaml

run-client:
	go run cmd/client/main.go