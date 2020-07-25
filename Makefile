run:
	protoc --proto_path=raw raw/*.proto --go_out=plugins=grpc:uploadpb

clean:
	rm uploadpb/*

up:
	go run main.go

test:
	go test -cover -race ./...