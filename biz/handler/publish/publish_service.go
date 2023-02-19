// Code generated by hertz generator.

package publish

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"simple_tiktok/biz/dal"
	"simple_tiktok/biz/model/feed"
	"simple_tiktok/biz/model/publish"
	"simple_tiktok/biz/model/user"
	"simple_tiktok/biz/redis"
	"simple_tiktok/pojo"
	"simple_tiktok/util"
	"strconv"
)

// PublishAction .
// @router /douyin/publish/action/ [POST]
func PublishAction(ctx context.Context, c *app.RequestContext) {
	var req publish.PublishRequest
	req.Title = c.PostForm("title")

	resp := new(publish.PublishResponse)
	message := ""
	resp.StatusMsg = &message
	username, exists := c.Get("user_name")
	if !exists {
		resp.StatusCode = 1
		message = "解析Token失败，没有Token解析的信息"
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

	video := pojo.Video{
		UserName:  username.(string),
		Title:     req.Title,
		VideoPath: req.Title + ".mp4",
		CoverPath: req.Title + ".jpg",
	}
	if err := c.SaveUploadedFile(file, "./data/"+video.VideoPath); err != nil {
		resp.StatusCode = 1
		message = "保存文件失败"
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}
	if err := util.Cover("./data/"+video.VideoPath, "./data/"+video.CoverPath); err != nil {
		resp.StatusCode = 1
		message = "获取视频封面失败"
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}
	// 执行数据库事务
	if err := dal.CreateVedio(&video); err != nil {
		resp.StatusCode = 1
		message = "数据库寄寄"
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}

	err = redis.InitLikeCount(int64(video.ID))
	if err != nil {
		resp.StatusCode = 1
		message = "redis寄寄"
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}

	resp.StatusCode = 0
	message = "success"
	c.JSON(consts.StatusOK, resp)
}

// GetPublishList .
// @router /douyin/publish/list/ [GET]
func GetPublishList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req publish.PublishListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(publish.PublishListResponse)
	message := ""
	resp.StatusMsg = &message
	username, exists := c.Get("user_name")
	if !exists {
		resp.StatusCode = 1
		message = "解析Token失败，没有Token解析的信息"
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	videos, err := dal.FindVideoByName(username.(string))
	if err != nil {
		resp.StatusCode = 1
		message = "没找到视频相关信息，寄"
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}

	resp.StatusCode = 0
	users, err := dal.FindUserByName(username.(string))
	if err != nil {
		resp.StatusCode = 1
		message = "没找到用户相关信息，寄"
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}
	if len(users) > 1 {
		resp.StatusCode = 1
		message = "找到的用户太多，寄"
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}

	// 假设你有一个名为videoSlice的存储了多个VideoInfo的切片

	var videoList []*feed.VideoInfo
	for _, v := range videos {
		//fmt.Println(v.CoverPath)
		like, err := redis.GetLikeCount(int64(v.ID))
		if err != nil {
			resp.StatusCode = 1
			message = "redis，寄"
			c.JSON(consts.StatusInternalServerError, resp)
		}
		likeInfo, err := redis.GetLikeInfo(strconv.Itoa(int(v.ID)), strconv.Itoa(int(users[0].ID)))
		if err != nil {
			resp.StatusCode = 1
			message = "Redis查询是否点赞信息出错"
			c.JSON(consts.StatusBadRequest, resp)
			return
		}

		var isLike bool
		if likeInfo == nil {
			isLike = false
		} else {
			isLike = true
		}
		video := &feed.VideoInfo{
			ID: int64(v.ID),
			Author: &user.UserInfo{
				ID:            int64(users[0].ID),
				Name:          users[0].UserName,
				FollowerCount: 0,
				IsFollow:      false,
			},
			PlayURL:       "http://192.168.137.1:8888/data/" + v.VideoPath,
			CoverURL:      "http://192.168.137.1:8888/data/" + v.CoverPath,
			FavoriteCount: like,
			CommentCount:  0,
			IsFavorite:    isLike,
			Title:         v.Title,
		}
		videoList = append(videoList, video)
	}

	resp.VideoList = videoList
	message = "success"
	c.JSON(consts.StatusOK, resp)
}
