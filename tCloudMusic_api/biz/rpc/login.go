package rpc

import (
	"context"
	"tCloudMusic_api/kitex_gen/tCloudMisc/biz/api/bizservice"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
)

type LoginRpc struct {
	client *bizservice.Client
}

func NewLoginRpc() *LoginRpc {
	c, err := bizservice.NewClient("biz", client.WithHostPorts("127.0.0.1:8889"))
	if err != nil {
		klog.Errorf("create LoginRpc error: %v\n", err.Error())
		return &LoginRpc{}
	}
	return &LoginRpc{client: &c}
}

// Login: user login rpc interfaces
func (l *LoginRpc) Login(ctx context.Context, username, password string) string {
	return ""
}

// Logout: user logout rpc interfaces
func (l *LoginRpc) Logout(ctx context.Context, username string) bool {
	return false
}
