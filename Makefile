
all: client server

dep:
	@echo "--> Running dep"
	@dep ensure

protoc:
	@echo "--> Generating Go files"
	cd  math/delivery/grpc/math_grpc && protoc --go_out=plugins=grpc:. math.proto

server:
	@echo "--> Building server"
	go build -i -o server cmd\server\main.go

client:
	@echo "--> Building client"
	go build -i -o client cmd\client\main.go

.PHONY: client server protoc dep
