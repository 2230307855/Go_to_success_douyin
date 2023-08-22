package dao

import (
	"douyin/config"
	"douyin/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// SetupDB 初始化数据库和 ORM
func SetupDB() {

	// 获取数据库配置
	config, err := config.GetConfig("db")
	if err != nil {
		panic("获取数据库配置失败")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.GetString("mysql.username"),
		config.GetString("mysql.password"),
		config.GetString("mysql.host"),
		config.GetInt("mysql.port"),
		config.GetString("mysql.dbname"))

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("链接数据库失败, error=" + err.Error())
	}

	db.AutoMigrate(&models.User{}, &models.Video{}, &models.Comment{}, &models.FavoriteVideoRelation{}, &models.FollowRelation{}, &models.Message{}, &models.FavoriteCommentRelation{})

}
