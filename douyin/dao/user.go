package dao

import (
	"douyin/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// 查询username是否存在
// 如果 查询到用户 返回 true
// 如果 没查询到用户 返回 false
func Checkname(username string) bool {
	var user models.User
	if err := db.Where(&models.User{UserName: username}).First(&user).Error; err != nil {
		return false
	}

	return true
}

// 创建新用户
func CreateUser(newUser *models.User) {
	db.Create(newUser)
}

// 登录判断密码
func CheckPassword(username, password string) (uint, bool) {
	var user models.User
	result := db.Where(&models.User{UserName: username}).First(&user)

	// 如果用户存在
	if result.Error == nil {
		// 用户存在，验证密码
		if user.Password == password {
			return user.ID, true
		} else {
			return 0, false
		}
	}

	return 0, false
}

// 用于用户信息
func CheckId(userid uint) (*models.User, error) {
	var user models.User
	result := db.Preload("FavoriteVideos.Author").Preload("FavoriteVideos").First(&user, userid)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, result.Error
	}

	return &user, nil
}

// 用户发布视频的时候，作品数加1
func WorkCountAdd(userId int) error {
	var user models.User
	result1 := db.First(&user, userId)
	if result1.Error != nil {
		fmt.Println("can't find user")
		return result1.Error
	}
	originWorkCount := user.WorkCount
	result2 := db.Model(&user).Update("work_count", originWorkCount+1)
	if result2.Error != nil {
		fmt.Println("can't add work count")
		return result2.Error
	}
	return nil
}

// 自己的关注总数加1\减去1，是否已经关注is_follow改为true
func AddAttentionUpdateIsFollow(userId int, opType int) error {
	var user models.User
	result1 := db.First(&user, userId)
	if result1.Error != nil {
		return result1.Error
	}
	if opType == 1 {
		resFollowingCount := int(user.FollowingCount + 1)
		result2 := db.Model(&user).Updates(models.User{
			FollowingCount: uint(resFollowingCount),
		})
		if result2.Error != nil {
			return result2.Error
		}
		return nil
	} else {
		resFollowingCount := int(user.FollowingCount - 1)
		//不受0值影响
		mp := map[string]interface{}{"FollowingCount": resFollowingCount}
		result3 := db.Model(&user).Updates(mp)
		if result3.Error != nil {
			return result3.Error
		}
		return nil
	}
}

// 粉丝数加1\减去1
func AddFollowerCount(userId int, opType int) error {
	var user models.User
	result1 := db.First(&user, userId)
	if result1.Error != nil {
		return result1.Error
	}
	if opType == 1 {
		resFlowerCount := int(user.FollowerCount + 1)
		result2 := db.Model(&user).Updates(models.User{
			FollowerCount: uint(resFlowerCount),
		})
		if result2.Error != nil {
			return result2.Error
		}
		return nil
	} else {
		resFlowerCount := int(user.FollowerCount - 1)
		//不受0值影响
		mp := map[string]interface{}{"FollowerCount": resFlowerCount}
		result3 := db.Model(&user).Updates(mp)
		if result3.Error != nil {
			return result3.Error
		}
		return nil
	}
}
