proto:
	protoc --proto_path=pkg/auth/pb --go_out=pkg/auth/pb --go_opt=paths=source_relative auth.proto

start:

build:
