generate-proto:
	protoc -I=./ --go_out=${GOPATH}/src ./internal/models/activity.proto --proto_path=../proto-example/include ;
	 protoc -I=./ --go_out=${GOPATH}/src ./internal/models/user.proto --proto_path=../proto-example/include ;
	 	protoc -I=./ --go_out=${GOPATH}/src ./internal/service/user_service.proto --go-grpc_out=${GOPATH}/src --proto_path=../proto-example/include ;
	 		protoc -I=./ --go_out=${GOPATH}/src ./internal/service/activity_service.proto --go-grpc_out=${GOPATH}/src --proto_path=../proto-example/include ;

inject-gorm:
	./scripts/inject_gorm.sh

server:
	go build ./cmd/server

client:
	go build ./cmd/client