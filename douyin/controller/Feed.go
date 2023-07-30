package controller

import (
	"douyin/models"
	"net/http"
	"time"

	"douyin/dao"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	models.Response
	NextTime  int64          `json:"next_time,omitempty"`
	VideoList []models.Video `json:"video_list,omitempty"`
}

func Feed(c *gin.Context) {

	videos, err := dao.GetAllVideos()
	if err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, FeedResponse{
		Response:  models.Response{StatusCode: 0},
		NextTime:  time.Now().Unix(),
		VideoList: videos,
	})
}
