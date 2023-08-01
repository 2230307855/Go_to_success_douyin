package controller

import (
	"douyin/dao"
	"douyin/models"
	"douyin/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FavaorteList struct {
	models.Response
	VideoList []VideoWithFavoriteStatus `json:"video_list"`
}

type VideoWithFavoriteStatus struct {
	models.Video
	IsFavorite bool `json:"is_favorite"`
}

// 登录用户对视频的点赞和取消点赞操作
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	video_id := c.Query("video_id")
	action_type := c.Query("action_type")
	// user_id := 1//

	// 校验 Token
	user_id, err1 := utils.GetIdFromToken(token)
	// 在获取token中id的时候校验失败
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg":  "login expired or illegal token",
		})
		return
	}

	// 将字符串转换int64格式 但最后需要uint格式
	v_id, errVideo := strconv.ParseUint(video_id, 10, 64)
	a_type, errUser := strconv.ParseUint(action_type, 10, 64)

	if errVideo != nil || errUser != nil {
		c.JSON(400, gin.H{"error": "Invalid ID or type"})
		return
	}

	// 进行点赞和取消点赞操作
	// 获取作者ID
	author_id, err := dao.VideoAuthorId(uint(v_id))

	if err != nil {
		c.JSON(http.StatusNotFound, models.Response{
			StatusCode: 1,
			StatusMsg:  "视频错误",
		})
	}

	// 如果是点赞
	if a_type == 1 {
		err := dao.CreateVideoFavorite(uint(user_id), author_id, uint(v_id))
		if err != nil {
			c.JSON(http.StatusBadRequest, models.Response{
				StatusCode: 1,
			})
			return
		}
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 0,
			StatusMsg:  "点赞成功",
		})
		return
	}

	if err := dao.DelFavoriteUserVideoId(uint(user_id), author_id, uint(v_id)); err != nil {
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		StatusCode: 0,
		StatusMsg:  "取消点赞成功",
	})

}

// 用户与视频的点赞关系
// 通过用户id和视频id
// func GetFavoriteVideoRelation()

// 用户的所有点赞视频
func FavoriteList(c *gin.Context) {

	token := c.Query("token")

	// 校验 Token
	user_id, err1 := utils.GetIdFromToken(token)
	// 在获取token中id的时候校验失败
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg":  "login expired or illegal token",
		})
		return
	}

	user, err := dao.CheckId(uint(user_id))
	if err != nil {
		c.JSON(http.StatusNotFound, models.Response{
			StatusCode: 1,
			StatusMsg:  "用户错误",
		})
		return
	}

	// 构建带有 is_favorite 字段的视频列表
	var videoListWithFavoriteStatus []VideoWithFavoriteStatus
	for _, video := range user.FavoriteVideos {
		videoListWithFavoriteStatus = append(videoListWithFavoriteStatus, VideoWithFavoriteStatus{
			Video:      video,
			IsFavorite: true, // 视频在用户的喜欢列表中，所以设为 true
		})
	}

	c.JSON(http.StatusOK, FavaorteList{
		Response:  models.Response{StatusCode: 0},
		VideoList: videoListWithFavoriteStatus,
	})

}
