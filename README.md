# dummy-storage-wrapper-communicator

A dummy Go program that talks to a storage bench wrapper.

## Regenerating .pb.gos

```
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u google.golang.org/grpc
protoc --go_out=plugins=grpc:proto/ *.proto
```