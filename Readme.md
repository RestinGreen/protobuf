# How to generate the Go files from .proto

1. cd proto-files
2. run this command

protoc --go_out=../generated --go_opt=paths=source_relative \
    --go-grpc_out=../generated --go-grpc_opt=paths=source_relative \
    ./dex.proto

3. cd ..
4. git add .
5. git commit -m "message"
6. git push