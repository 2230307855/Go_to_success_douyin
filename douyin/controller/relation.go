package controller

import (
	"douyin/dao"
	"douyin/models"
	"douyin/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FansListResponse struct {
	models.Response
	FansList []models.User `json:"user_list"`
}

type FriendListResponse struct {
	models.Response
	FriendList []models.User `json:"user_list"`
}

// 关注操作
func RelationAction(c *gin.Context) {

}

// 关注列表
func FollowList(c *gin.Context) {

}

// 粉丝列表
func FollowerList(c *gin.Context) {
	user_id := c.Query("user_id")
	token := c.Query("token")

	// 校验 Token
	_, err1 := utils.GetIdFromToken(token)
	// 在获取token中id的时候校验失败
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg":  "login expired or illegal token",
		})
		return
	}

	// 将字符串转换int64格式 但最后需要uint格式
	userID, errID := strconv.ParseUint(user_id, 10, 64)
	if errID != nil {
		c.JSON(400, gin.H{"error": "Invalid ID or type"})
		return
	}

	fans, err := dao.GetFollowingListByUserID(uint(userID))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: 1,
			StatusMsg:  "查询粉丝错误",
		})
		return
	}

	c.JSON(http.StatusOK, FansListResponse{
		Response: models.Response{StatusCode: 0, StatusMsg: "Success"},
		FansList: fans,
	})
}

// 好友列表
func FriendList(c *gin.Context) {
	user_id := c.Query("user_id")
	token := c.Query("token")

	// 校验 Token
	_, err1 := utils.GetIdFromToken(token)
	// 在获取token中id的时候校验失败
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg":  "login expired or illegal token",
		})
		return
	}

	// 将字符串转换int64格式 但最后需要uint格式
	userID, errID := strconv.ParseUint(user_id, 10, 64)
	if errID != nil {
		c.JSON(400, gin.H{"error": "Invalid ID or type"})
		return
	}

	friends, err := dao.GetFriendListByUserID(uint(userID))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: 1,
			StatusMsg:  "查询朋友错误",
		})
		return
	}

	c.JSON(http.StatusOK, FriendListResponse{
		Response:   models.Response{StatusCode: 0, StatusMsg: "Success"},
		FriendList: friends,
	})
}
