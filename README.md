# task-manager-api-gateway

## Description

This is a golang backend api gateway microservice with a REST server related to a Task Management API. This is created using gin framework. Api gateway receives request from clients as REST API calls and directs them to the relevant microservice via gRPC calls. This api gateway acts as the gRPC client. 

## Folder structure

```
├───api
│   └───docs
├───cmd
└───pkg
    └───auth
        ├───handlers
        └───pb
```

- api - contains the REST API specifications/documentation using swagger
- cmd - contains the main application entrypoint
- pkg - contains application core
 - auth - contains logic relevant to authentication microservice
    - handlers - handlers of the authentication microservice
    - pb - proto file and generated protobuf files

## Prerequisites
- Golang installed in your system (this project use go version 1.19)
- Docker installed in your system
- Make installed in your system

# Getting started

1. Clone the repository

```
$ git clone https://github.com/niluwats/task-manager-api-gateway.git
```

2. Navigate to the project directory

```
$ cd task-manager-api-gateway
```

3. Build and run application using make

```
$ make up_build
```

4. Access the application

Once the containers are up and running, you can access the application at http://localhost:8080


## Dependencies
- [gin](github.com/gin-gonic/gin) - web framework for REST server
- [go-grpc](google.golang.org/grpc) - gRPC library to set up client
- [gin-swagger](github.com/swaggo/gin-swagger) - gin swagger library to for API specification

```
$ go get github.com/gin-gonic/gin
$ go get google.golang.org/grpc
$ go get github.com/swaggo/gin-swagger
$ go get github.com/swaggo/files
```
