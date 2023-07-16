proto:
	protoc --proto_path=pkg/auth/pb --go_out=pkg/auth/pb --go_opt=paths=source_relative auth.proto

up_build:
	docker-compose down
	docker-compose up --build

up:
	docker-compose down
	docker-compose up

down:
	docker-compose down
