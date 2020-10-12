# GO-PROTO-USER_ACTIVITY

## Repository structure

```
.
├── cmd
│   ├── client
│   │   └── main.go
│   └── server
│       └── main.go
├── config
│   ├── config.go
│   └── init.go
├── config.yaml
├── go.mod
├── go.sum
├── internal
│   ├── models
│   │   ├── activity.proto
│   │   └── user.proto
│   └── service
│       ├── activity_service.proto
│       └── user_service.proto
├── Makefile
├── scripts
│   ├── gorm.sh
│   └── inject_gorm.sh
└── src
    ├── model
    │   ├── activity.pb.go
    │   └── user.pb.go
    ├── service
    │   ├── activity_service_grpc.pb.go
    │   ├── activity_service.pb.go
    │   ├── user_service_grpc.pb.go
    │   └── user_service.pb.go
    └── store
        ├── activity_store.go
        └── user_store.go
```

## Internal
Internal contains all the proto files. separeted in to models and services.
The model folder contains models proto file for user an activity.

To generate go files from proto, use the following make command
    
    make generate-proto

This would generate .pb.go files in sr/model nad src/service folders respectively

## Model
Contains generated .pb.go files for model proto files

## Service
Contains generated .pb.go files for service proto files

## Store
Store implements services for bot user and activity

the user service has following methods
```go
type UserServiceClient interface {
	Add(ctx context.Context, in *model.User, opts ...grpc.CallOption) (*UserResponse, error)
	Get(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserResponse, error)
}
```

the activity service has following methods

```go
	Add(ctx context.Context, in *model.Activity, opts ...grpc.CallOption) (*ActivityResponse, error)
	Get(ctx context.Context, in *GetActivityRequest, opts ...grpc.CallOption) (*ActivityResponse, error)
	Start(ctx context.Context, in *GetActivityRequest, opts ...grpc.CallOption) (*ActivityResponse, error)
	Stop(ctx context.Context, in *GetActivityRequest, opts ...grpc.CallOption) (*ActivityResponse, error)
	IsValid(ctx context.Context, in *GetActivityRequest, opts ...grpc.CallOption) (*BoolResult, error)
	IsDone(ctx context.Context, in *GetActivityRequest, opts ...grpc.CallOption) (*BoolResult, error)
```

The activity can be of the following types

    “play”, “sleep”, “eat” and “read”

and implements 

    Start, Stop, IsValid and IsDone

Note: instead of implementing the original problem (play, sleep,eat and read)
the new implementations meant more sense to me

## How to run Steps

- Generate go files from protofiles run `make generate-proto`
- Inject gorm meta tags `make inject-gorm`
- RUN `go mod tidy` to build dependencies
-  build server and client `make server` `make client`

To run the server and client , 
    execute server 

    ./server

and then client binaries

    ./client
