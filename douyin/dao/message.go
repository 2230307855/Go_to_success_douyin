package dao

import (
	"douyin/models"
	"time"
)

//在数据库添加一条message
func AddMessage(message models.Message) ( error){
	err :=db.Create(&message).Error
	if err != nil {
		return err
	}
	return err
}

//根据userID返回Message列表
//创建一个键值对，健为token，值为时间戳
func GetMessageListByUserID(fromID uint, toID uint, timeStam int64)([]models.Message, error){
	//查询所有fromid与toid之间的message
	var messageList []models.Message
	err := db.Model(&models.Message{}).Where("(created_at > ?) AND ((from_user_id = ? AND to_user_id = ?) OR (from_user_id = ? AND to_user_id = ?))",
	time.Unix(timeStam,0), fromID, toID, toID, fromID).Order("created_at ASC").Find(&messageList).Error
	if err != nil{
		return nil, err 
	}
	return  messageList,nil	
}