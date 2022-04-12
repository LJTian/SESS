package handler

import (
	"context"
	"fmt"
	"testing"

	UserProto "SESS/api/rpc/proto"
	"SESS/cmd/grpc_svr_user/global"
)

func TestUserServer_GetUserList(t *testing.T) {

	InitDB()
	var userServer UserServer
	global.DB = DB

	list, err := userServer.GetUserList(context.Background(), &UserProto.PageInfo{
		Pn:    0,
		PSize: 10,
	})
	if err != nil {
		return
	}
	fmt.Println(list)
}

func TestUserServer_CreatUser(t *testing.T) {

	InitDB()
	var userServer UserServer
	global.DB = DB

	list, err := userServer.CreateUser(context.Background(), &UserProto.CreateUserInfo{
		NickName: "田利军1",
		Mobile:   "17611231237",
		PassWord: "123456",
	})
	if err != nil {
		return
	}
	fmt.Println(list)
}

func TestUserServer_GetUserByMobile(t *testing.T) {

	InitDB()
	var userServer UserServer
	global.DB = DB

	list, err := userServer.GetUserByMobile(context.Background(), &UserProto.MobileRequest{
		Mobile: "17611231236",
	})
	if err != nil {
		return
	}
	fmt.Println(list)
}

func TestUserServer_GetUserById(t *testing.T) {

	InitDB()
	var userServer UserServer
	global.DB = DB

	list, err := userServer.GetUserById(context.Background(), &UserProto.IdRequest{
		Id: 2,
	})
	if err != nil {
		return
	}
	fmt.Println(list)
}

func TestUserServer_UpdateUser(t *testing.T) {

	InitDB()
	var userServer UserServer
	global.DB = DB

	list, err := userServer.UpdateUser(context.Background(), &UserProto.UpdateUserInfo{
		Id:       2,
		NickName: "田利军",
		Gender:   "0",
	})
	if err != nil {
		return
	}
	fmt.Println(list)
}

func TestUserServer_CheckPassWord(t *testing.T) {

	InitDB()
	var userServer UserServer
	global.DB = DB

	list, err := userServer.CheckPassWord(context.Background(), &UserProto.PasswordCheckInfo{
		Password:          "123456",
		EncryptedPassword: "$pbkdf2-sha512$133dCp0dFEGPt2dj$3141ffacfaf6d28aa02257de74f0cc897cb7293d939c3fe51ff5c974f521feed",
	})
	if err != nil {
		return
	}
	fmt.Println(list)
}
