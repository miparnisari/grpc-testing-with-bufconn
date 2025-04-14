# grpc-testing-with-bufconn

## Setup

1. Install Go.
2. Install `buf`: https://buf.build/docs/installation/
3. Run

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```