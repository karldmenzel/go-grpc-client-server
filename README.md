# Go gRPC Client-Server Demo

```bash
go run client/main/client_main.go
```

```bash
go run server/main/server_main.go
```

```bash
go test ./...
```

Compile the protocol buffer:
```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative shared/math_server.proto
```