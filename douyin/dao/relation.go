package dao

import (
	"douyin/models"
	"errors"
)

// 返回粉丝列表
// 要在relations表中查询被关注的用户(to_user_id)，然后通过对应关注的用户找到对应的
func GetFollowingListByUserID(userID uint) ([]models.User, error) {

	var fansRelations []models.FollowRelation

	// 查询所有粉丝的FollowRelation结构
	if err := db.Model(&models.FollowRelation{}).Where("to_user_id = ?", userID).Find(&fansRelations).Error; err != nil {
		return nil, err
	}

	// 如果没有粉丝
	if len(fansRelations) == 0 {
		return nil, errors.New("no fans found")
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

	var followings []models.FollowRelation

	// 1，查询用户（user_id）关注的FollowRelation结构
	if err := db.Model(&models.FollowRelation{}).Where("user_id = ?", userID).Find(&followings).Error; err != nil {
		return nil, err
	}

	// 如果没有关注的人就直接返回了
	if len(followings) == 0 {
		return nil, errors.New("没有关注的人")
	}

	// 提取用户关注的的 UserID
	var followingUserIDs []uint
	for _, following := range followings {
		followingUserIDs = append(followingUserIDs, following.ToUserID)
	}

	// 2，查询关注用户的粉丝 (to_user_id)
	var fansRelations []models.FollowRelation
	if err := db.Model(&models.FollowRelation{}).Where("to_user_id = ?", userID).Find(&fansRelations).Error; err != nil {
		return nil, err
	}

	// 如果用户一个粉丝也没有
	if len(fansRelations) == 0 {
		return nil, errors.New("没有粉丝")
	}

	// 提取用户粉丝的 UserID
	var fansUserIDs []uint
	for _, fans := range fansRelations {
		fansUserIDs = append(fansUserIDs, fans.UserID)
	}

	// 3，找出互相关注的用户ID，才是朋友
	var friendUserIDs []uint
	for _, userID := range followingUserIDs {
		for _, fansUserID := range fansUserIDs {
			if userID == fansUserID {
				friendUserIDs = append(friendUserIDs, userID)
			}
		}
	}

	// 4，更具3找出的朋友 ID 找出信息
	var friends []models.User
	if err := db.Model(&models.User{}).Find(&friends, friendUserIDs).Error; err != nil {
		return nil, err
	}

	return friends, nil

}

// 在relation中添加或删除一条关注记录
func AddRelation(userId int, targetUserId, opType int) error {
	newReCord := models.FollowRelation{UserID: uint(userId), ToUserID: uint(targetUserId)}
	if opType == 1 {
		result := db.Create(&newReCord)
		if result.Error != nil {
			return result.Error
		} else {
			return nil
		}
	} else {
		//先检索出来
		db.First(&newReCord)
		//再删除
		result2 := db.Unscoped().Delete(&models.FollowRelation{}, newReCord.ID)
		if result2.Error != nil {
			return result2.Error
		} else {
			return nil
		}
	}

}

// 根据用户的id查询出其关注的用户列表信息
func GetAttentionUserById(userId int) ([]models.User, error) {
	var relations []models.FollowRelation
	userList := make([]models.User, 0)
	result := db.Preload("ToUser").Find(&relations, "user_id=?", userId)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, relation := range relations {
		userList = append(userList, relation.ToUser)
	}
	return userList, nil
}

// 用于判断两者之间是否有关注的关系
func IsHaveRelation(user_id int, to_user_id int) bool {
	var relationObj models.FollowRelation
	result := db.Where("user_id = ? and to_user_id = ?", user_id, to_user_id).First(&relationObj)
	if result.RowsAffected < 1 {
		return false
	}
	if result.Error != nil {
		return false
	}
	return true
}
