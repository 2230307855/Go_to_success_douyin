package dao

import (
	"douyin/models"
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
func GetMessageListByUserID(fromID uint, toID uint)([]models.Message, error){
	//查询所有fromid与toid之间的message
	var messageList []models.Message
	err := db.Model(&models.Message{}).Where("from_user_id = ? AND to_user_id = ?", fromID, toID).Or("from_user_id = ? AND to_user_id = ?", toID, fromID).Find(&messageList).Error
	if err != nil{
		return nil, err 
	}

	return  messageList,nil	
}