package dao

import (
	"douyin/models"
)

// 返回粉丝列表
// 要在relations表中查询被关注的用户(to_user_id)，然后通过对应关注的用户找到对应的
func GetFollowingListByUserID(userID uint) ([]models.User, error) {

	var fansRelations []models.FollowRelation

	// 查询所有粉丝的FollowRelation结构
	if err := db.Model(&models.FollowRelation{}).Where("to_user_id = ?", userID).Find(&fansRelations).Error; err != nil {
		return nil, err
	}

	// 提取所有粉丝的 UserID
	var fansUserIDs []uint
	for _, relation := range fansRelations {
		fansUserIDs = append(fansUserIDs, relation.UserID)
	}

	// 查询所有粉丝用户信息
	var fans []models.User
	if err := db.Model(&models.User{}).Find(&fans, fansUserIDs).Error; err != nil {
		return nil, err
	}

	return fans, nil

}

// 返回朋友列表
func GetFriendListByUserID(userID uint) ([]models.User, error) {

	var friendRelations []models.FollowRelation

	// 查询所有朋友的FollowRelation结构
	if err := db.Model(&models.FollowRelation{}).Where("user_id = ?", userID).Find(&friendRelations).Error; err != nil {
		return nil, err
	}

	// 提取所有朋友的的 UserID
	var friendUserIDs []uint
	for _, relation := range friendRelations {
		friendUserIDs = append(friendUserIDs, relation.ToUserID)
	}

	// 查询所有朋友用户信息
	var friends []models.User
	if err := db.Model(&models.User{}).Find(&friends, friendUserIDs).Error; err != nil {
		return nil, err
	}

	return friends, nil

}
