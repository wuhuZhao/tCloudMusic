package main

import (
	"context"
	"tCloudMusic_rpc/biz/common"
	"tCloudMusic_rpc/biz/dao"
	"tCloudMusic_rpc/biz/model"
	api "tCloudMusic_rpc/kitex_gen/tCloudMisc/biz/api"

	"gorm.io/gorm"
)

// BizServiceImpl implements the last service interface defined in the IDL.
type BizServiceImpl struct {
	userDao *dao.UserDao
}

func NewBizServiceImpl(db *gorm.DB) *BizServiceImpl {
	return &BizServiceImpl{userDao: dao.NewUserDao(db)}
}

// Login implements the BizServiceImpl interface.
func (s *BizServiceImpl) Login(ctx context.Context, request *api.LoginRequest) (resp *api.LoginReponse, err error) {
	// TODO: Your code here...
	// return
	user := s.userDao.FindUserByUserNameAndPassword(request.GetUsername(), request.GetPassword())
	if user.Password != request.GetPassword() || user.Username != request.GetUsername() {
		resp.SetResp(&api.Response{Code: -1, Msg: "user not exist"})
		resp.SetToken("")
		return
	}
	token, err := common.GenToken(request.Username)
	if err != nil {
		resp.SetResp(&api.Response{Code: -1, Msg: "gen token error"})
		resp.SetToken("")
		return
	}
	//TODO: add redis token
	resp.SetResp(&api.Response{Code: 1, Msg: "user exist"})
	resp.SetToken(token)
	return
}

// Logout implements the BizServiceImpl interface.
func (s *BizServiceImpl) Logout(ctx context.Context, request *api.LogoutRequest) (resp *api.LogoutResponse, err error) {
	// TODO: Your code here...
	// TODO: add redis token to logout
	resp.SetResult_(true)
	resp.SetResp(&api.Response{Code: 1, Msg: "logout success"})
	return
}

// SignUp implements the BizServiceImpl interface.
func (s *BizServiceImpl) SignUp(ctx context.Context, request *api.SignUpRequest) (resp *api.SignUpResponse, err error) {
	user := model.User{Username: request.GetUsername(), Email: request.GetEmail(), Password: request.GetPassword()}
	_, err = s.userDao.InsertUser(&user)
	resp.SetResult_(true)
	if err != nil {
		resp.SetResult_(false)
	}
	return
}
