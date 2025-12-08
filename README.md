# Go gRPC Client-Server Demo

This project is a demo of gRPC in the Go programming language.

The server exposes four functions for doing math operations, 
and four functions for getting the total count of invocations of each math function.

The client makes 1000 calls to the server, each for a random math function.
It then gets the count of how many times each function has been called. 

The shared functions and objects are defined in [./magicMath/magic_math.proto](./magicMath/magic_math.proto).

## Commands

To start the server:
```bash
go run client/main/client_main.go
```

To run a client:
```bash
go run server/main/server_main.go
```

To run the unit tests:
```bash
go test ./...
```

To recompile the protocol buffer files:
```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative magicMath/magic_math.proto
```