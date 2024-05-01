generate:
	protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.

run-server:
	go run main.go
	
run-client:
	grpcui --plaintext 127.0.0.1:8080
