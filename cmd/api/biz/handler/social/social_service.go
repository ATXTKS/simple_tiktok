// Code generated by hertz generator.

package social

import (
	"context"
	"simple_tiktok/kitex_gen/social"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	Modelsocial "simple_tiktok/cmd/api/biz/model/social"
	RpcSocial "simple_tiktok/cmd/api/biz/rpc/social"
)

// FollowAction .
// @router /douyin/relation/action/ [POST]
func FollowAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req Modelsocial.FollowRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	userID, _ := c.Get("userid")
	ID := userID.(int64)
	userName, _ := c.Get("username")
	Name := userName.(string)
	resp, err := RpcSocial.FollowAction(ctx, &social.FollowRequest{
		Token:      req.Token,
		ToUserId:   req.ToUserID,
		ActionType: req.ActionType,
		UserId:     &ID,
		UserName:   &Name,
	})
	if err != nil {
		resp = new(social.FollowResponse)
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	c.JSON(consts.StatusOK, resp)
}

// GetFollowList .
// @router /douyin/relation/follow/list/ [GET]
func GetFollowList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req Modelsocial.FollowingListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp, err := RpcSocial.GetFollowList(ctx, &social.FollowingListRequest{
		UserId: req.UserID,
		Token:  req.Token,
	})
	if err != nil {
		resp = new(social.FollowingListResponse)
		resp.StatusCode = 1
		msg := err.Error()
		resp.StatusMsg = &msg
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	c.JSON(consts.StatusOK, resp)
}

// GetFollowerList .
// @router /douyin/relation/follower/list/ [GET]
func GetFollowerList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req Modelsocial.FollowerListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	userName, _ := c.Get("username")
	Name := userName.(string)
	resp, err := RpcSocial.GetFollowerList(ctx, &social.FollowerListRequest{
		UserId:   req.UserID,
		Token:    req.Token,
		UserName: &Name,
	})
	if err != nil {
		resp = new(social.FollowerListResponse)
		resp.StatusCode = 1
		msg := err.Error()
		resp.StatusMsg = &msg
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	c.JSON(consts.StatusOK, resp)
}

// GetFriendList .
// @router /douyin/relation/friend/list/ [GET]
func GetFriendList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req Modelsocial.FriendListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp, err := RpcSocial.GetFriendList(ctx, &social.FriendListRequest{
		UserId: req.UserID,
		Token:  req.Token,
	})
	if err != nil {
		resp = new(social.FriendListResponse)
		resp.StatusCode = 1
		msg := err.Error()
		resp.StatusMsg = &msg
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	c.JSON(consts.StatusOK, resp)
}

// MessageChat .
// @router /douyin/message/chat/ [GET]
func MessageChat(ctx context.Context, c *app.RequestContext) {
	var err error
	var req Modelsocial.MessageChatReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	userID, _ := c.Get("userid")
	ID := userID.(int64)
	userName, _ := c.Get("username")
	Name := userName.(string)
	resp, err := RpcSocial.MessageChat(ctx, &social.MessageChatReq{
		Token:      req.Token,
		ToUserId:   req.ToUserID,
		PreMsgTime: req.PreMsgTime,
		UserId:     &ID,
		UserName:   &Name,
	})
	if err != nil {
		resp = new(social.MessageChatResp)
		resp.StatusCode = 1
		msg := err.Error()
		resp.StatusMsg = &msg
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	c.JSON(consts.StatusOK, resp)
}

// MessageAction .
// @router /douyin/message/action/ [POST]
func MessageAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req Modelsocial.MessageActionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	userID, _ := c.Get("userid")
	ID := userID.(int64)
	userName, _ := c.Get("username")
	Name := userName.(string)
	resp, err := RpcSocial.MessageAction(ctx, &social.MessageActionReq{
		Token:      req.Token,
		ToUserId:   req.ToUserID,
		ActionType: req.ActionType,
		Content:    req.Content,
		UserId:     &ID,
		UserName:   &Name,
	})
	if err != nil {
		resp = new(social.MessageActionResp)
		resp.StatusCode = 1
		msg := err.Error()
		resp.StatusMsg = &msg
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	c.JSON(consts.StatusOK, resp)
}
