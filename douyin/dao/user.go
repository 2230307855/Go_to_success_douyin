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
