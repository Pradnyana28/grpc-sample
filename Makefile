run: ## generate protobuf file to go file
	protoc --proto_path=raw raw/*.proto --go_out=plugins=grpc:uploadpb

clean: ## clean the generated protobuf
	rm uploadpb/*

server: ## start server
	go run cmd/server/main.go -port 5050

client: ## start client server
	go run cmd/client/main.go -address 0.0.0.0:5050

test: ## run unit testing
	go test -cover -race ./...