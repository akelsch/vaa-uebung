# vaa-uebung

## Build

Run `go build` or better `go install` in any `/cmd` subdirectory to install an executable needed to run the scripts.

## Protocol Buffers Support

1. Install protoc (v3.19.2), e.g. using `scoop install protobuf`
2. Install protoc-gen-go (v1.27.1), e.g. using `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`

Compilation of proto files can be done using `protoc --proto_path=./api/pb --go_out=./api/pb --go_opt=paths=source_relative ./api/pb/*.proto` relative to one of the submodules.

## IntelliJ Support

1. Install [Go plugin](https://plugins.jetbrains.com/plugin/9568-go)
2. Ensure Go modules integration is enabled in `File | Settings | Languages & Frameworks | Go | Go Modules`
3. Add proto files to Protocol Buffers import path in `File | Settings | Languages & Frameworks | Protocol Buffers`
    - e.g. `C:/Users/Arthur/git/vaa-uebung/ueb01/api/pb`
