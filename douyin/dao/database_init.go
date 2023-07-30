package dao

import (
	"fmt"

	"douyin/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// SetupDB 初始化数据库和 ORM
func SetupDB() {

	// dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
	// 	viper.GetString("database.username"),
	// 	viper.GetString("database.password"),
	// 	viper.GetString("database.host"),
	// 	viper.GetInt("database.port"),
	// 	viper.GetString("database.database"),
	// 	viper.GetString("database.charset"),
	// )
	//----------------------------------------------------------
	//发布时候的配置
	username := "datatest"
	password := "EYxwxRsNfYTn7SN7"
	host := "111.92.243.152"
	port := 3306
	dbname := "datatest"
	//本地运行配置------------------------------------------------
	//username := "root"
	//password := "123456"
	//host := "127.0.0.1"
	//port := 3306
	//dbname := "gorm_test"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("链接数据库失败, error=" + err.Error())
	}

	// 将 gorm 设计的表映射到Mysql数据库中
	db.AutoMigrate(&models.User{}, &models.Video{}, &models.Comment{}, &models.FavoriteCommentRelation{}, &models.FavoriteVideoRelation{}, &models.FollowRelation{},
		&models.Comment{})

}
