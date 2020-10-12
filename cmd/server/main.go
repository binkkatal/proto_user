package main

import (
	"fmt"
	"log"
	"net"

	"github.com/binkkatal/proto_user/config"
	"github.com/binkkatal/proto_user/src/model"
	"github.com/binkkatal/proto_user/src/service"
	"github.com/binkkatal/proto_user/src/store"

	"google.golang.org/grpc"
)

func main() {

	db := config.GetDB()
	netListener := getNetListener(7000)
	grpcServer := grpc.NewServer()

	err := db.AutoMigrate(
		&model.User{},
		&model.Activity{},
	)
	if err != nil {
		log.Fatalf("failed automigrate", err)
	}
	userStore := store.NewUserStore(db)
	activityStore := store.NewActivityStore(db)
	service.RegisterUserServiceServer(grpcServer, userStore)
	service.RegisterActivityServiceServer(grpcServer, activityStore)

	// start the server
	if err := grpcServer.Serve(netListener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}

func getNetListener(port uint) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(fmt.Sprintf("failed to listen: %v", err))
	}

	return lis
}
