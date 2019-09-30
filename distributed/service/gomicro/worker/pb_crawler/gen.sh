
protoc --proto_path=$GOPATH/src:. --micro_out=$GOPATH/src --go_out=$GOPATH/src *.proto
