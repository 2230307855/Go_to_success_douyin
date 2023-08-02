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
	token := c.Query("token")
	targetUserId := c.Query("to_user_id")
	action_type := c.Query("action_type")
	optionType, _ := strconv.Atoi(action_type)
	id, err := utils.GetIdFromToken(token)
	//token校验失败
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg":  "login expired or illegal token",
		})
		return
	}
	//token校验成功
	//自己的关注总数加1
	err1 := dao.AddAttentionUpdateIsFollow(id, optionType)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg":  "option error",
		})
		return
	}
	//被关注者的粉丝数加1
	tarId, _ := strconv.Atoi(targetUserId)
	err2 := dao.AddFollowerCount(tarId, optionType)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg":  "option error",
		})
		return
	}
	//在关系表中添加一条记录，内容填充的有关注着与被关注着的信息
	err3 := dao.AddRelation(id, tarId, optionType)
	if err3 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg":  "option error",
		})
		return
	}
	//关注成功
	if optionType == 1 {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  "关注成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  "取关成功",
		})
	}
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
