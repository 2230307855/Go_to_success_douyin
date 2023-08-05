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

type UserLoginResponse struct {
	models.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	models.Response
	User models.User `json:"user"`
}

// 注册
func Register(c *gin.Context) {

	username := c.Query("username")
	password := c.Query("password")
	if check := dao.Checkname(username); check {
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 1,
			StatusMsg:  "User already exist",
		})

	} else {
		newUser := models.User{
			UserName: username,
			Password: password,
			Avatar:   "https://web-test-store.oss-cn-hangzhou.aliyuncs.com/douyin/avatars/default.jpg",
		}
		dao.CreateUser(&newUser)
		id, _ := dao.CheckPassword(username, password)
		authToken := utils.JwtGeneration(username, int(id))
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: models.Response{StatusCode: 0},
			UserId:   int64(newUser.ID),
			Token:    authToken,
		})
	}

}

// 登录
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	if check := dao.Checkname(username); !check {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: models.Response{StatusCode: 1, StatusMsg: "User does not exist"},
		})
		return
	}
	id, exist := dao.CheckPassword(username, password)

	// 如果密码不对应
	if !exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: models.Response{StatusCode: 1, StatusMsg: "password error"},
		})
		return
	}

	token := utils.JwtGeneration(username, int(id))
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: models.Response{StatusCode: 0},
		UserId:   int64(id),
		Token:    token,
	})

}

// 用户信息
func UserInfo(c *gin.Context) {
	user_id := c.Query("user_id")
	token := c.Query("token")
	u, err := strconv.ParseUint(user_id, 10, 64)
	//不合法的用户id
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
		return
	}
	//用户不存在
	user, err := dao.CheckId(uint(u))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	//用户token过期
	if err := utils.TokenVerify(token, int(u)); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user login expired or illegal status"})
		return
	}
	//正确的用户信息
	c.JSON(http.StatusOK, UserResponse{
		Response: models.Response{StatusCode: 0},
		User:     *user,
	})

}
