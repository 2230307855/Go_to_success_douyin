package controller

import (
	"douyin/dao"
	"douyin/models"
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

	token := username + password

	if check := dao.Checkname(username); check {
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 1,
			StatusMsg:  "User already exist",
		})

	} else {
		newUser := models.User{
			UserName: username,
			Password: password,
		}
		dao.CreateUser(&newUser)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: models.Response{StatusCode: 0},
			UserId:   int64(newUser.ID),
			Token:    token,
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

	c.JSON(http.StatusOK, UserLoginResponse{
		Response: models.Response{StatusCode: 0},
		UserId:   int64(id),
		Token:    username + password,
	})

}

// 用户信息
func UserInfo(c *gin.Context) {
	user_id := c.Query("user_id")

	u, err := strconv.ParseUint(user_id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
		return
	}

	user, err := dao.CheckId(uint(u))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, UserResponse{
		Response: models.Response{StatusCode: 0},
		User:     *user,
	})

}
