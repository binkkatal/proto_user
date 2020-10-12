package main

import (
	"log"

	"github.com/binkkatal/proto_user/src/model"
	"github.com/binkkatal/proto_user/src/service"

	"context"
	"fmt"

	"google.golang.org/grpc"
)

func main() {
	serverAddress := "localhost:7000"

	conn, e := grpc.Dial(serverAddress, grpc.WithInsecure())

	if e != nil {
		panic(e)
	}
	defer conn.Close()

	userClient := service.NewUserServiceClient(conn)

	userModel := model.User{
		Name:        "test-demo",
		Email:       "cc@cc.com",
		PhoneNumber: "123",
	}

	if responseMessage, e := userClient.Add(context.Background(), &userModel); e != nil {
		panic(fmt.Sprintf("Was not able to insert Record %v", e))
	} else {
		fmt.Println("Record Inserted..")
		fmt.Println(responseMessage)
		fmt.Println("=============================")
	}

	ctx := context.Background()
	userResp, err := userClient.Get(ctx, &service.GetUserRequest{Id: userModel.Id})
	if err != nil {
		log.Fatalf("Failed to get user %s", err)
	}
	user := userResp.User

	activityClient := service.NewActivityServiceClient(conn)
	activity := &model.Activity{
		ActivityType: model.Activity_PLAY,
		UserId:       user.Id,
	}
	_, err = activityClient.Add(ctx, activity)
	if err != nil {
		log.Fatalf("Failed to fetch activity %s", err.Error())
	}
	actReq := &service.GetActivityRequest{Id: activity.Id}
	_, err = activityClient.Start(ctx, actReq)
	if err != nil {
		log.Fatalf("Failed to start activity %s", err.Error())
	}

	validResp, err := activityClient.IsValid(ctx, actReq)
	if err != nil {
		log.Fatalf("Failed to start activity %s", err.Error())
	}
	if validResp.Result {
		fmt.Println("Activity is valid")
	} else {
		fmt.Println("Activity is invalid")
	}
	_, err = activityClient.Stop(ctx, actReq)
	if err != nil {
		log.Fatalf("Failed to start activity %s", err.Error())
	}
	validResp, err = activityClient.IsValid(ctx, actReq)
	if err != nil {
		log.Fatalf("Failed to start activity %s", err.Error())
	}
	if validResp.Result {
		fmt.Println("Activity is valid")
	} else {
		fmt.Println("Activity is invalid")
	}

}
