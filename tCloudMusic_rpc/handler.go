package main

import (
	"context"
	"tCloudMusic_rpc/biz/dao"
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
	resp.SetResp(&api.Response{Code: 1, Msg: "user exist"})
	resp.SetToken(user.Password)
	return
}

// Logout implements the BizServiceImpl interface.
func (s *BizServiceImpl) Logout(ctx context.Context, request *api.LogoutRequest) (resp *api.LogoutResponse, err error) {
	// TODO: Your code here...
	resp.SetResult_(true)
	resp.SetResp(&api.Response{Code: 1, Msg: "logout success"})
	return
}
