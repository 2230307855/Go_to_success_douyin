package controller

import (
	"douyin/dao"
	"douyin/models"
	"douyin/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RelationResponse struct {
	models.Response
	UserList []RelationWithFollow `json:"user_list"`
}

type RelationWithFollow struct {
	models.User
	Isfollow bool `json:"is_follow"`
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
	//自己的关注总数加1或者减去1
	err1 := dao.AddAttentionUpdateIsFollow(id, optionType)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg":  "option error",
		})
		return
	}
	//被关注者的粉丝数加1或者减去1
	tarId, _ := strconv.Atoi(targetUserId)
	err2 := dao.AddFollowerCount(tarId, optionType)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg":  "option error",
		})
		return
	}
	//在关系表中添加或删除一条记录，内容填充的有关注着与被关注着的信息
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
	//获取参数
	token := c.Query("token")
	userId := c.Query("user_id")
	//token校验失败
	myId, _ := strconv.Atoi(userId)
	err := utils.TokenVerify(token, myId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg":  "Invalid login,please login again",
		})
		return
	}
	//获取关注列表失败
	attentionUserList, err2 := dao.GetAttentionUserById(myId)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg":  "option error",
		})
		return
	}
	relationWithFollows := make([]RelationWithFollow, 0)
	for _, val := range attentionUserList {
		//从关系表中查找两用户之间是否有关注记录
		isFollow := dao.IsHaveRelation(myId, int(val.ID))
		fmt.Println("is_follow is:", isFollow)
		relationWithFollows = append(relationWithFollows, RelationWithFollow{
			User:     val,
			Isfollow: isFollow,
		})
	}
	//给了用户的id，查询其关注的用户的列表成功
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "success",
		"user_list":   relationWithFollows,
	})
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

	// 构建带有 is_follow 字段用户列表
	var relationWithFollow []RelationWithFollow
	for _, user := range fans {
		relationWithFollow = append(relationWithFollow, RelationWithFollow{
			User:     user,
			Isfollow: false,
		})
	}

	c.JSON(http.StatusOK, RelationResponse{
		Response: models.Response{StatusCode: 0, StatusMsg: "Success"},
		UserList: relationWithFollow,
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

	// 构建带有 is_follow 字段用户列表
	var relationWithFollow []RelationWithFollow
	for _, user := range friends {
		relationWithFollow = append(relationWithFollow, RelationWithFollow{
			User:     user,
			Isfollow: true,
		})
	}

	c.JSON(http.StatusOK, RelationResponse{
		Response: models.Response{StatusCode: 0, StatusMsg: "Success"},
		UserList: relationWithFollow,
	})
}
