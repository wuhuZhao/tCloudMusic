package rpc

import (
	"context"
	"tCloudMusic_api/kitex_gen/tCloudMisc/biz/api"
	"tCloudMusic_api/kitex_gen/tCloudMisc/biz/api/bizservice"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
)

type LoginRpc struct {
	client bizservice.Client
}

func NewLoginRpc() *LoginRpc {
	c, err := bizservice.NewClient("biz", client.WithHostPorts("127.0.0.1:8889"))
	if err != nil {
		klog.Errorf("create LoginRpc error: %v\n", err.Error())
		return &LoginRpc{}
	}
	return &LoginRpc{client: c}
}

// Login: user login rpc interfaces
func (l *LoginRpc) Login(ctx context.Context, username, password string) string {
	resp, err := l.client.Login(ctx, &api.LoginRequest{Username: username, Password: password})
	klog.Infof("[login] resp: %v\n", resp)
	if err != nil {
		klog.Errorf("[login] get err from kitex server, err: %v\n", err.Error())
		return ""
	}
	return resp.GetToken()
}

// Logout: user logout rpc interfaces
func (l *LoginRpc) Logout(ctx context.Context, username string) bool {
	resp, err := l.client.Logout(ctx, &api.LogoutRequest{Username: username})
	klog.Infof("[logout] resp: %v\n", resp)
	if err != nil {
		klog.Errorf("[logout] get err from kitex server, err: %v\n", err.Error())
		return false
	}
	return resp.GetResult_()
}
