// Code generated by hertz generator.

package base

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	ModelBase "simple_tiktok/cmd/api/biz/model/base"
	RpcBase "simple_tiktok/cmd/api/biz/rpc/base"
	"simple_tiktok/kitex_gen/base"
	"simple_tiktok/util/jwt"
)

// UserRegister .
// @router /douyin/user/register/ [POST]
func UserRegister(ctx context.Context, c *app.RequestContext) {
	var err error
	var req ModelBase.RegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp, err := RpcBase.CreateUser(ctx, &base.RegisterRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		resp = new(base.RegisterResponse)
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	c.JSON(consts.StatusOK, resp)
}

// UserLogin .
// @router /douyin/user/login/ [POST]
func UserLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req ModelBase.LoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp, err := RpcBase.UserLogin(ctx, &base.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		resp = new(base.LoginResponse)
		resp.StatusCode = 1
		errMsg := err.Error()
		resp.StatusMsg = &errMsg
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	c.JSON(consts.StatusOK, resp)
}

// GetUserInfo .
// @router /douyin/user/ [GET]
func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req ModelBase.UserInfoRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp, err := RpcBase.GetUserInfo(ctx, &base.UserInfoRequest{
		UserId: req.UserID,
		Token:  req.Token,
	})
	if err != nil {
		resp = new(base.UserInfoResponse)
		resp.StatusCode = 1
		errMsg := err.Error()
		resp.StatusMsg = &errMsg
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	c.JSON(consts.StatusOK, resp)
}

// GetVideoList .
// @router /douyin/feed/ [GET]
func GetVideoList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req ModelBase.FeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	if req.Token != nil {
		mc, _ := jwt.ParseToken(*req.Token)
		req.UserID = &mc.UserID
	}
	resp, err := RpcBase.GetVideoList(ctx, &base.FeedRequest{
		LatestTime: req.LatestTime,
		UserId:     req.UserID,
	})
	if err != nil {
		resp = new(base.FeedResponse)
		resp.StatusCode = 1
		errMsg := err.Error()
		resp.StatusMsg = &errMsg
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// PublishAction .
// @router /douyin/publish/action/ [POST]
func PublishAction(ctx context.Context, c *app.RequestContext) {
	//var err error
	//var req ModelBase.PublishRequest
	//err = c.BindAndValidate(&req)
	//if err != nil {
	//	c.String(consts.StatusBadRequest, err.Error())
	//	return
	//}
	resp := new(base.PublishResponse)
	message := ""
	resp.StatusMsg = &message

	title := c.PostForm("title")
	if title == "" {
		resp.StatusCode = 1
		message = "没有标题！"
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	file, err := c.FormFile("data")
	//fmt.Println(file)
	if err != nil {
		resp.StatusCode = 1
		message = "文件上传失败"
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	// 这里拿到视频，先保存，保存的位置是相对于当前程序运行位置的
	if err := c.SaveUploadedFile(file, "../../data/"+title+".mp4"); err != nil {
		resp.StatusCode = 1
		message = "保存文件失败"
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}

	fmt.Println("===============================")
	fmt.Println(c.Get("userid"))
	fmt.Println(c.Get("username"))
	fmt.Println("===============================")
	userID, _ := c.Get("userid")
	userName, _ := c.Get("username")
	resp, err = RpcBase.PublishAction(ctx, &base.PublishRequest{
		Data:     nil,
		Token:    "",
		Title:    title,
		UserId:   userID.(int64),
		UserName: userName.(string),
	})
	if err != nil {
		resp = new(base.PublishResponse)
		resp.StatusCode = 1
		errMsg := err.Error()
		resp.StatusMsg = &errMsg
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	c.JSON(consts.StatusOK, resp)
}

// GetPublishList .
// @router /douyin/publish/list/ [GET]
func GetPublishList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req ModelBase.PublishListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp, err := RpcBase.GetPublishList(ctx, &base.PublishListRequest{
		UserId:  req.UserID,
		QueryId: req.QueryID,
	})
	if err != nil {
		resp = new(base.PublishListResponse)
		resp.StatusCode = 1
		errMsg := err.Error()
		resp.StatusMsg = &errMsg
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	c.JSON(consts.StatusOK, resp)
}
