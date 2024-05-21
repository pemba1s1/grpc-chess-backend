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
