// Code generated by Kitex v0.10.1. DO NOT EDIT.

package relationservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	relation "github.com/mutezebra/tiktok/pkg/kitex_gen/api/relation"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Follow(ctx context.Context, req *relation.FollowReq, callOptions ...callopt.Option) (r *relation.FollowResp, err error)
	GetFollowList(ctx context.Context, req *relation.GetFollowListReq, callOptions ...callopt.Option) (r *relation.GetFollowListResp, err error)
	GetFansList(ctx context.Context, req *relation.GetFansListReq, callOptions ...callopt.Option) (r *relation.GetFansListResp, err error)
	GetFriendsList(ctx context.Context, req *relation.GetFriendsListReq, callOptions ...callopt.Option) (r *relation.GetFriendsListResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kRelationServiceClient{
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

type kRelationServiceClient struct {
	*kClient
}

func (p *kRelationServiceClient) Follow(ctx context.Context, req *relation.FollowReq, callOptions ...callopt.Option) (r *relation.FollowResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Follow(ctx, req)
}

func (p *kRelationServiceClient) GetFollowList(ctx context.Context, req *relation.GetFollowListReq, callOptions ...callopt.Option) (r *relation.GetFollowListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFollowList(ctx, req)
}

func (p *kRelationServiceClient) GetFansList(ctx context.Context, req *relation.GetFansListReq, callOptions ...callopt.Option) (r *relation.GetFansListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFansList(ctx, req)
}

func (p *kRelationServiceClient) GetFriendsList(ctx context.Context, req *relation.GetFriendsListReq, callOptions ...callopt.Option) (r *relation.GetFriendsListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFriendsList(ctx, req)
}
