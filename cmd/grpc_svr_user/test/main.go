package main

import (
	UserProto "SESS/api/rpc/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"
)

func main() {

	conn, err := grpc.Dial("192.168.124.5:49494", grpc.WithInsecure())
	if err != nil {
		return
	}

	defer conn.Close()

	client := UserProto.NewUserGrpcServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := client.CreateUser(ctx, &UserProto.CreateUserInfo{
		NickName: "18612121213",
		PassWord: "123456",
		Mobile:   "18612121215",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(user)

}
