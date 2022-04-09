package handler

import (
	UserProto "SESS/api/rpc/proto"
	"context"
	"go.uber.org/zap"
)

type UserServer struct {
}

func (receiver UserServer) CreatUser(context.Context,
	*UserProto.CreatUserInfo) (*UserProto.UserInfoResponse, error) {

	zap.S().Info("调用[CreatUser] 成功")
	return nil, nil
}
