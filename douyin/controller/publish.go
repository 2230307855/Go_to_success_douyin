package controller

import (
	"douyin/dao"
	"douyin/models"
	"douyin/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// 登录用户选择视频上传
func Publish(c *gin.Context) {
	token := c.Request.FormValue("token")
	title := c.Request.FormValue("title")
	//先校验token
	id, err1 := utils.GetIdFromToken(token)
	//在获取token中id的时候校验失败
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg":  "login expired or illegal token",
		})
		return
	}
	//用户登录合法，将视频上传到oss服务器并将发布信息存入数据库
	filePath, filename := utils.SingleFileUploadMidWare(c)
	url := utils.UploadVideo(filePath, filename)
	err := dao.PublishVideo(models.Video{
		CreatedAt:     time.Now().UTC(),
		UpdatedAt:     time.Now().UTC(),
		AuthorID:      uint(id),
		PlayUrl:       url,
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         title,
		CoverUrl:      "https://web-test-store.oss-cn-hangzhou.aliyuncs.com/douyin/covers/default.png",
	})
	err2 := dao.WorkCountAdd(id)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg":  "work_count add filed please check it",
		})
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg":  "upload success but database upload failed please check it",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status_code": http.StatusOK,
			"status_msg":  "文件上传成功",
		})
	}
}

// 用户的视频发布列表，直接列出用户所有投稿过的视频
func PublishList(c *gin.Context) {
	token := c.Query("token")
	id := c.Query("user_id")
	myId, _ := strconv.Atoi(id)
	//校验token是否合法
	err := utils.TokenVerify(token, myId)
	//登录不合法
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg":  "login expired or illegal",
			"video_list":  "{}",
		})
		return
	}
	//登录合法
	var author models.AuthorOfVideo
	dao.GetAuthorById(myId, &author)
	//make([]models.VideoItem, 0)
	videoList := []models.VideoItem{}
	videoListres := make([]models.VideoItem, 0)
	dao.GetVideoListByAuthorId(myId, &videoList)

	for _, v := range videoList {
		v.Author = author
		videoListres = append(videoListres, v)
	}
	resObj := models.GetVideoListResponse{
		StatusCode: http.StatusOK,
		StatusMsg:  "请求成功！",
		VideoList:  videoListres,
	}
	c.JSON(http.StatusOK, resObj)
}
