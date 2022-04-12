package handler

import (
	"context"
	"crypto/sha512"
	"fmt"
	"strings"
	"time"

	"github.com/anaskhan96/go-password-encoder"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"

	UserProto "SESS/api/rpc/proto"
	"SESS/cmd/grpc_svr_user/global"
	"SESS/cmd/grpc_svr_user/model"
)

type UserServer struct {
}

func ModelToRsponse(user model.User) UserProto.UserInfoResponse {
	//在grpc的message中字段有默认值，你不能随便赋值nil进去，容易出错
	//这里要搞清， 哪些字段是有默认值
	userInfoRsp := UserProto.UserInfoResponse{
		Id:       int32(user.ID),
		PassWord: user.PassWord,
		NickName: user.NickName,
		Gender:   user.Gender,
		Role:     user.Role,
		Mobile:   user.Mobile,
	}
	if user.Birthday != nil {
		userInfoRsp.BirthDay = uint64(user.Birthday.Unix())
	}
	return userInfoRsp
}

func Paginate(page, pageSize uint32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(int(offset)).Limit(int(pageSize))
	}
}

// GetUserByMobile 根据手机号获取用户信息
func (receiver UserServer) GetUserByMobile(ctx context.Context,
	request *UserProto.MobileRequest) (*UserProto.UserInfoResponse, error) {

	userInfo := model.User{}
	resp := global.DB.Where(&model.User{Mobile: request.Mobile}).First(&userInfo)
	if resp.Error != nil {
		zap.S().Errorf("GetUserByMobile err is [%s]", resp.Error.Error())
		return nil, status.Errorf(codes.Internal, "查找用户失败")
	}
	respUserInfo := ModelToRsponse(userInfo)
	return &respUserInfo, nil
}

// GetUserById 根据ID获取用户信息
func (receiver UserServer) GetUserById(ctx context.Context,
	request *UserProto.IdRequest) (*UserProto.UserInfoResponse, error) {

	userInfo := model.User{}
	resp := global.DB.First(&userInfo, request.Id)
	if resp.Error != nil {
		zap.S().Errorf("GetUserByMobile err is [%s]", resp.Error.Error())
		return nil, status.Errorf(codes.Internal, "查找用户失败")
	}
	respUserInfo := ModelToRsponse(userInfo)
	return &respUserInfo, nil
}

// CreateUser 创建用户
func (receiver UserServer) CreateUser(ctx context.Context,
	info *UserProto.CreateUserInfo) (*UserProto.UserInfoResponse, error) {

	user := model.User{
		Mobile:   info.Mobile,
		NickName: info.NickName,
	}
	// 构建表信息
	global.DB.AutoMigrate(&model.User{})

	//密码加密
	options := &password.Options{16, 100, 32, sha512.New}
	salt, encodedPwd := password.Encode(info.PassWord, options)
	user.PassWord = fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)

	resp := global.DB.Create(&user)
	if resp.Error != nil {
		zap.S().Errorf("CreatUser err is [%s]", resp.Error.Error())
		return nil, status.Errorf(codes.Internal, "填加用户失败")
	}

	respUser := ModelToRsponse(user)
	zap.S().Info("调用[CreatUser] 成功")
	return &respUser, nil
}

// UpdateUser 更新用户
func (receiver UserServer) UpdateUser(ctx context.Context,
	info *UserProto.UpdateUserInfo) (*emptypb.Empty, error) {

	user := model.User{}

	resp := global.DB.First(&user, info.Id)
	if resp.RowsAffected == 0 {
		zap.S().Errorf("查找用户信息失败")
		return nil, status.Errorf(codes.NotFound, "查找用户信息失败")
	}

	user.Gender = info.Gender
	user.NickName = info.NickName
	birthDay := time.Unix(int64(info.BirthDay), 0)
	user.Birthday = &birthDay

	resp = global.DB.Save(&user)
	if resp.Error != nil {
		zap.S().Errorf("UpdateUser err is [%s]", resp.Error.Error())
		return nil, status.Errorf(codes.Internal, "更改用户信息失败")
	}

	return &emptypb.Empty{}, nil
}

// CheckPassWord 校验密码
func (receiver UserServer) CheckPassWord(ctx context.Context,
	info *UserProto.PasswordCheckInfo) (*UserProto.CheckResponse, error) {

	//校验密码
	options := &password.Options{16, 100, 32, sha512.New}
	passwordInfo := strings.Split(info.EncryptedPassword, "$")
	check := password.Verify(info.Password, passwordInfo[2], passwordInfo[3], options)
	return &UserProto.CheckResponse{Success: check}, nil
}

// GetUserList 获取用户列表
func (receiver UserServer) GetUserList(ctx context.Context,
	info *UserProto.PageInfo) (*UserProto.UserListResponse, error) {

	var users []model.User
	resp := global.DB.Find(&users)
	if resp.Error != nil {
		zap.S().Errorf("GetUserList err is [%s]", resp.Error.Error())
		return nil, status.Errorf(codes.Internal, "查找用户列表失败")
	}

	var respUserList UserProto.UserListResponse
	respUserList.Total = int32(resp.RowsAffected)

	global.DB.Scopes(Paginate(info.Pn, info.PSize)).Find(&users)
	zap.S().Info(users)
	for _, k := range users {
		respUserInfo := ModelToRsponse(k)
		respUserList.Data = append(respUserList.Data, &respUserInfo)
	}
	return &respUserList, nil
}
