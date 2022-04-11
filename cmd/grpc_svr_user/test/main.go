package main

import (
	UserProto "SESS/api/rpc/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"
)

func main() {

	conn, err := grpc.Dial("127.0.0.1:62029", grpc.WithInsecure())
	if err != nil {
		return
	}

	defer conn.Close()

	client := UserProto.NewUserGrpcServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := client.CreatUser(ctx, &UserProto.CreatUserInfo{
		NickName: "18612121212",
		PassWord: "123456",
		Mobile:   "18612121212",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(user)

}
