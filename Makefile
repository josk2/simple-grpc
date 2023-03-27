.PHONY: gen_protoc_helloworld install_grpc_package


install_grpc_package:
	go get -u google.golang.org/grpc

gen_protoc_helloworld:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        helloworld/protoc/greeter.proto