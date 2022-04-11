package handler

import (
	UserProto "SESS/api/rpc/proto"
	"context"
	"fmt"
	"go.uber.org/zap"
)

type UserServer struct {
}

func (receiver UserServer) CreatUser(ctx context.Context,
	userInfo *UserProto.CreatUserInfo) (*UserProto.UserInfoResponse, error) {

	zap.S().Info("调用[CreatUser] 成功")
	fmt.Println("调用[CreatUser] 成功")
	return &UserProto.UserInfoResponse{
		UserId:   0,
		NickName: userInfo.NickName,
		PassWord: userInfo.PassWord,
		Mobile:   userInfo.Mobile,
	}, nil
}
