This is a chess game server made in go using grpc
## Prerequisites

- Go. For installation instructions, see Go’s [Getting Started](https://golang.org/doc/install) guide.
- Protocol buffer complier. For installation instructions, see [Protocol Buffer Compiler Installation](https://grpc.io/docs/protoc-installation/).
- Go Plugins for protocol compiler.
    1. Install the protocol compiler plugins for Go using the following commands: 
    ```bash
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
    ```
    2. Update your PATH so that the protoc compiler can find the plugins:
    ```bash
    export PATH="$PATH:$(go env GOPATH)/bin"
    ```
- MakeFile. For installation instructions, see [Make for Windows](https://gnuwin32.sourceforge.net/packages/make.htm)

## Local Development
### Clone Repo
```bash
git clone https://github.com/pemba1s1/chess-backend.git
```
### Generate the gRPC client and server interfaces from our .proto service definition
```bash
make generate_grpc_code
```
### Start Envoy
As browser can only communicate using HTTP/1 but gRPC uses HTTP/2. So we need to use envoy to catch HTTP/1 request from browser and point to corresponding gRPC service.
```bash
docker-compose up
```

### Build and run docker
```bash
make build_docker
make run_docker
```