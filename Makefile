run: ## generate protobuf file to go file
	protoc --proto_path=raw raw/*.proto --go_out=plugins=grpc:uploadpb

clean: ## clean the generated protobuf
	rm uploadpb/*

up: ## start server
	go run main.go

test: ## run unit testing
	go test -cover -race ./...