package dao

import (
	"douyin/models"
	"fmt"

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
	// 发布配置---------------------------------------------
	username := "gorm_test"
	password := "NbxxFrRi8BNsctxn"
	host := "8.130.82.173"
	port := 3306
	dbname := "gorm_test"
	//登录配置-ChenglongShi---------------------------------
	//username := "gorm_test"
	//password := "TYnkmaGnDzrXJciw"
	//host := "111.92.243.152"
	//port := 3306
	//dbname := "gorm_test"

	// 数据库配置-JintongXu
	// username := "demo"
	// password := "n5fpLARZE46EsBPA"
	// host := "8.130.82.173"
	// port := 3306
	// dbname := "demo"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("链接数据库失败, error=" + err.Error())
	}

	db.AutoMigrate(&models.User{}, &models.Video{}, &models.Comment{}, &models.FavoriteVideoRelation{}, &models.FollowRelation{}, &models.Message{}, &models.FavoriteCommentRelation{})

}
