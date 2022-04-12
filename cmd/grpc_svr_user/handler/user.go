package handler

import (
	UserProto "SESS/api/rpc/proto"
	"SESS/cmd/grpc_svr_user/global"
	"SESS/cmd/grpc_svr_user/model"
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct {
}

// CreatUser 创建用户
func (receiver UserServer) CreatUser(ctx context.Context,
	userInfo *UserProto.CreatUserInfo) (*UserProto.UserInfoResponse, error) {

	user := model.User{
		Mobile:   userInfo.Mobile,
		NickName: userInfo.NickName,
		PassWord: userInfo.PassWord,
	}

	// 构建表信息
	global.DB.AutoMigrate(&model.User{})

	resp := global.DB.Create(&user)
	if resp.Error != nil {
		zap.S().Errorf("CreatUser err is [%s]", resp.Error.Error())
		return nil, status.Errorf(codes.Internal, "填加用户失败")
	}

	respUser := UserProto.UserInfoResponse{}
	respUser.UserId = int32(user.ID)
	respUser.NickName = user.NickName
	respUser.PassWord = user.PassWord
	respUser.Mobile = user.Mobile

	zap.S().Info("调用[CreatUser] 成功")
	return &respUser, nil
}
