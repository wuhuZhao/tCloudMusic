// Code generated by Kitex v0.4.4. DO NOT EDIT.

package bizservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	api "tCloudMusic_rpc/kitex_gen/tCloudMisc/biz/api"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Login(ctx context.Context, request *api.LoginRequest, callOptions ...callopt.Option) (r *api.LoginReponse, err error)
	Logout(ctx context.Context, request *api.LogoutRequest, callOptions ...callopt.Option) (r *api.LogoutResponse, err error)
	SignUp(ctx context.Context, request *api.SignUpRequest, callOptions ...callopt.Option) (r *api.SignUpResponse, err error)
	SearchMusics(ctx context.Context, request *api.GetMusicListRequest, callOptions ...callopt.Option) (r *api.GetMusicListResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kBizServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kBizServiceClient struct {
	*kClient
}

func (p *kBizServiceClient) Login(ctx context.Context, request *api.LoginRequest, callOptions ...callopt.Option) (r *api.LoginReponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Login(ctx, request)
}

func (p *kBizServiceClient) Logout(ctx context.Context, request *api.LogoutRequest, callOptions ...callopt.Option) (r *api.LogoutResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Logout(ctx, request)
}

func (p *kBizServiceClient) SignUp(ctx context.Context, request *api.SignUpRequest, callOptions ...callopt.Option) (r *api.SignUpResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SignUp(ctx, request)
}

func (p *kBizServiceClient) SearchMusics(ctx context.Context, request *api.GetMusicListRequest, callOptions ...callopt.Option) (r *api.GetMusicListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SearchMusics(ctx, request)
}
