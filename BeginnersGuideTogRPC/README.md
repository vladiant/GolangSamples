# Beginners Guide to gRPC in Go

Based on [Beginners Guide to gRPC in Go!](https://www.youtube.com/watch?v=BdzYdN_Zd9Q) from [TutorialEdge](https://www.youtube.com/channel/UCwFl9Y49sWChrddQTD9QhRA) YouTube channel.

## Environment
* Ubuntu 20
* Visual Studio Code + [Go](https://marketplace.visualstudio.com/items?itemName=golang.go) plugin
* Go version 1.13.8

## Steps to create
1. Read to follow: [gRPC Quick start](https://grpc.io/docs/languages/go/quickstart/)
2. Install protocol compiler: `sudo apt install protobuf-compiler`
3. `export PATH="$PATH:$(go env GOPATH)/bin"`
4. `export GO111MODULE=on`
5. `go get google.golang.org/protobuf/cmd/protoc-gen-go google.golang.org/grpc/cmd/protoc-gen-go-grpc`
6. `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc`
7. Prepare files `server.go` , `chat.proto` and `chat` folder.
8. `protoc --go_out=./chat --go_opt=paths=source_relative --go-grpc_out=./chat --go-grpc_opt=paths=source_relative chat.proto`
9. Prepare `chat.go`
10. `go mod init chat` to create local module. Replaced `go mod init github.com/tutorialedge/go-grpc-tutorial/chat` from tutorial.
11. Prepare `client.go`
