package controller

import (
	"douyin/dao"
	"douyin/models"
	"douyin/utils"
	"net/http"
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
)

type MessageActionResponse struct{
	models.Response
	MessageList []models.Message `json:"message_list"`
}

// 发送信息
func MessageAction(c *gin.Context) {
	//获取信息
	token := c.Query("token")
	toUserIDStr := c.Query("to_user_id")
	actionType:= c.Query("action_type")
	content := c.Query("content")

	//验证token
	fromUserID, err := utils.GetIdFromToken(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg":  "login expired or illegal token",
		})
		return
	}

	//将to_user_id转为int64格式 但最后需要uint格式
	toUserID, err := strconv.ParseUint(toUserIDStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID or type"})
		return
	}

	//状态为1，发送消息
	if actionType == "1" {
		sendMess := models.Message{
			CreatedAt:  time.Now().UTC(),
			UpdatedAt:  time.Now().UTC(),
			FromUserID: uint(fromUserID),
			ToUserID: uint(toUserID),
			Content: content,
		}
		err := dao.AddMessage(sendMess)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.Response{
				StatusCode: 1,
				StatusMsg:  "Add information error",
			})
			return
		}
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 0,
			StatusMsg: "Message sent successfully",
		})
	}
}

// 聊天记录
func MessageChat(c *gin.Context) {
	//获取信息
	toUserIDStr := c.Query("to_user_id")
	token := c.Query("token")

	//校验token
	fromUserID, err := utils.GetIdFromToken(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg":  "Login expired or illegal token",
		})
		return
	}
	
	// 将字符串转换int64格式 但最后需要uint格式
	toUserID, err := strconv.ParseUint(toUserIDStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID or type"})
		return
	}

	//查询信息
	messages, err := dao.GetMessageListByUserID(uint(fromUserID) , uint(toUserID))
	if(err != nil){
		c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: 1,
			StatusMsg:  "Information query error",
		})
		return
	}

	c.JSON(http.StatusOK, MessageActionResponse{
		Response:   models.Response{
			StatusCode: 0, 
			StatusMsg: "Information query successfully",
		},
		MessageList: messages,
	})
}
