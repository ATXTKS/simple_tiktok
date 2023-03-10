// Code generated by Kitex v0.4.4. DO NOT EDIT.

package baseservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	base "simple_tiktok/kitex_gen/base"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	UserRegister(ctx context.Context, request *base.RegisterRequest, callOptions ...callopt.Option) (r *base.RegisterResponse, err error)
	UserLogin(ctx context.Context, request *base.LoginRequest, callOptions ...callopt.Option) (r *base.LoginResponse, err error)
	GetUserInfo(ctx context.Context, request *base.UserInfoRequest, callOptions ...callopt.Option) (r *base.UserInfoResponse, err error)
	GetVideoList(ctx context.Context, request *base.FeedRequest, callOptions ...callopt.Option) (r *base.FeedResponse, err error)
	PublishAction(ctx context.Context, request *base.PublishRequest, callOptions ...callopt.Option) (r *base.PublishResponse, err error)
	GetPublishList(ctx context.Context, request *base.PublishListRequest, callOptions ...callopt.Option) (r *base.PublishListResponse, err error)
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
	return &kBaseServiceClient{
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

type kBaseServiceClient struct {
	*kClient
}

func (p *kBaseServiceClient) UserRegister(ctx context.Context, request *base.RegisterRequest, callOptions ...callopt.Option) (r *base.RegisterResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserRegister(ctx, request)
}

func (p *kBaseServiceClient) UserLogin(ctx context.Context, request *base.LoginRequest, callOptions ...callopt.Option) (r *base.LoginResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserLogin(ctx, request)
}

func (p *kBaseServiceClient) GetUserInfo(ctx context.Context, request *base.UserInfoRequest, callOptions ...callopt.Option) (r *base.UserInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserInfo(ctx, request)
}

func (p *kBaseServiceClient) GetVideoList(ctx context.Context, request *base.FeedRequest, callOptions ...callopt.Option) (r *base.FeedResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetVideoList(ctx, request)
}

func (p *kBaseServiceClient) PublishAction(ctx context.Context, request *base.PublishRequest, callOptions ...callopt.Option) (r *base.PublishResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PublishAction(ctx, request)
}

func (p *kBaseServiceClient) GetPublishList(ctx context.Context, request *base.PublishListRequest, callOptions ...callopt.Option) (r *base.PublishListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetPublishList(ctx, request)
}
