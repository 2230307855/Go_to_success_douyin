package controller

import (
	"douyin/models"
	"douyin/utils"
	"net/http"
	"time"

	"douyin/dao"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	models.Response
	NextTime  int64          `json:"next_time,omitempty"`
	VideoList []models.Video `json:"video_list,omitempty"`
	Userid    int            `json:"userid"`
}

func Feed(c *gin.Context) {
	token := c.Query("token")

	user_id, err := utils.GetIdFromToken(token)

	videos, err := dao.GetAllVideos(user_id)
	if err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, FeedResponse{
		Response:  models.Response{StatusCode: 0},
		NextTime:  time.Now().Unix(),
		VideoList: videos,
	})
}
