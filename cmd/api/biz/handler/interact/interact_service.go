// Code generated by hertz generator.

package interact

import (
	"context"
	"simple_tiktok/kitex_gen/interact"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	ModelInteract "simple_tiktok/cmd/api/biz/model/interact"
	RpcInteract "simple_tiktok/cmd/api/biz/rpc/interact"
)

// LikeAction .
// @router /douyin/favorite/action/ [POST]
func LikeAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req ModelInteract.LikeRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	userID, _ := c.Get("userid")
	ID := userID.(int64)
	userName, _ := c.Get("username")
	Name := userName.(string)
	resp, err := RpcInteract.LikeAction(ctx, &interact.LikeRequest{
		Token:      req.Token,
		VideoId:    req.VideoID,
		ActionType: req.ActionType,
		UserId:     &ID,
		UserName:   &Name,
	})
	if err != nil {
		resp = new(interact.LikeResponse)
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	c.JSON(consts.StatusOK, resp)
}

// GetLikeList .
// @router /douyin/favorite/list/ [GET]
func GetLikeList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req ModelInteract.LikeListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp, err := RpcInteract.GetLikeList(ctx, &interact.LikeListRequest{
		UserId: req.UserID,
		Token:  req.Token,
	})
	if err != nil {
		resp = new(interact.LikeListResponse)
		resp.StatusCode = 1
		msg := err.Error()
		resp.StatusMsg = &msg
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	c.JSON(consts.StatusOK, resp)
}

// CommentAction .
// @router /douyin/comment/action/ [POST]
func CommentAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req ModelInteract.CommentRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	userID, _ := c.Get("userid")
	ID := userID.(int64)
	userName, _ := c.Get("username")
	Name := userName.(string)
	resp, err := RpcInteract.CommentAction(ctx, &interact.CommentRequest{
		Token:       req.Token,
		VideoId:     req.VideoID,
		ActionType:  req.ActionType,
		CommentText: req.CommentText, // ?????????????????????????????????action_type=1???????????????
		CommentId:   req.CommentID,   //??????????????????id??????action_type=2???????????????
		UserId:      &ID,
		UserName:    &Name,
	})
	if err != nil {
		resp = new(interact.CommentResponse)
		resp.StatusCode = 1
		msg := err.Error()
		resp.StatusMsg = &msg
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	c.JSON(consts.StatusOK, resp)
}

// GetCommentList .
// @router /douyin/comment/list/ [GET]
func GetCommentList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req ModelInteract.CommentListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp, err := RpcInteract.GetCommentList(ctx, &interact.CommentListRequest{
		Token:   req.Token,
		VideoId: req.VideoID,
	})
	if err != nil {
		resp = new(interact.CommentListResponse)
		resp.StatusCode = 1
		msg := err.Error()
		resp.StatusMsg = &msg
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	c.JSON(consts.StatusOK, resp)
}
