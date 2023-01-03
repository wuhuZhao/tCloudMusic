package login

import (
	"context"
	"tCloudMusic_api/biz/common"
	"tCloudMusic_api/biz/rpc"

	"github.com/cloudwego/hertz/pkg/app"
)

type loginHandler struct {
	rpc *rpc.LoginRpc
}

func NewLoginHandler() *loginHandler {
	return &loginHandler{rpc: rpc.NewLoginRpc()}
}

// Login: login api to users
func (l *loginHandler) Login(ctx context.Context, c *app.RequestContext) {
	username, password := c.PostForm("username"), c.PostForm("password")
	if len(username) == 0 || len(password) == 0 {
		c.JSON(-1, common.NewErrorResponse("username or password can't be empty, Please check it"))
	}
	c.JSON(0, common.NewSuccessResponse(l.rpc.Login(ctx, username, common.HashPassword(password))))
}

// Logout: logout api to users
func (l *loginHandler) Logout(ctx context.Context, c *app.RequestContext) {
	username := c.PostForm("username")
	if len(username) == 0 {
		c.JSON(-1, common.NewErrorResponse("username can't be empty, Please check it"))
	}
	c.JSON(0, common.NewSuccessResponse(l.rpc.Logout(ctx, username)))
}

// SignUp: SignUp api to users
func (l *loginHandler) SignUp(ctx context.Context, c *app.RequestContext) {
	username, password, email := c.PostForm("username"), c.PostForm("password"), c.PostForm("email")
	if len(username) == 0 || len(password) == 0 || len(email) == 0 {
		c.JSON(-1, common.NewErrorResponse("username or passwrod or email can't be empty, Please check it"))
	}
	c.JSON(0, common.NewSuccessResponse(l.rpc.SignUp(ctx, username, common.HashPassword(password), email)))
}
