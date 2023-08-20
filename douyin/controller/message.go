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

//传出去的CreatedAt类型为uint，所以新建类型intMessage
type intMessage struct{
	ID         uint      	  `json:"id"`
	ToUserID   uint           `json:"to_user_id"`
	FromUserID uint           `json:"from_user_id"`
	Content    string         `json:"content"`
	CreatedAt  uint 		  `json:"create_time"`
}

type MessageActionResponse struct{
	models.Response
	MessageList []intMessage `json:"message_list"`
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
	preMsgTime , err := strconv.ParseInt(c.Query("pre_msg_time"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg":  "pre_msg_time is not number",
		})
		return
	}
	
	// 若为非常大的time(认为是前端的bug)，则直接返回空消息
	
	if(preMsgTime >= 253402300799){
		var msglist []intMessage
		c.JSON(http.StatusOK, MessageActionResponse{
			Response:   models.Response{
				StatusCode: 0, 
				StatusMsg: "Information query successfully",
			},
			MessageList: msglist,
		})
		return 
	}

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
	messages, err := dao.GetMessageListByUserID(uint(fromUserID) , uint(toUserID), preMsgTime + 1)
	if(err != nil){
		c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: 1,
			StatusMsg:  "Information query error",
		})
		return
	}

	intMessageList := make([]intMessage, len(messages))
	for i, msg := range messages{
		intMessageList[i].ID = msg.ID
		intMessageList[i].Content = msg.Content
		intMessageList[i].FromUserID = msg.FromUserID
		intMessageList[i].ToUserID = msg.ToUserID
		intMessageList[i].CreatedAt = uint(msg.CreatedAt.Unix())
	}

	c.JSON(http.StatusOK, MessageActionResponse{
		Response:   models.Response{
			StatusCode: 0, 
			StatusMsg: "Information query successfully",
		},
		MessageList: intMessageList,
	})
}
