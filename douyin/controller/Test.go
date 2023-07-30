package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Test(c *gin.Context) {
	// username := viper.GetString("database.username")
	// password := viper.GetString("database.password")
	// host := viper.GetString("database.host")
	// port := viper.GetInt("database.port")
	// dbname := viper.GetString("database.dbname")

	username := "datatest"
	password := "EYxwxRsNfYTn7SN7"
	host := "111.92.243.152"
	port := 3306
	Dbname := "datatest"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)

	c.JSON(http.StatusOK, gin.H{
		"username": viper.GetString("database.username"),
		"password": viper.GetString("database.password"),
		"host":     viper.GetString("database.host"),
		"port":     viper.GetInt("database.port"),
		"dbname":   viper.GetString("database.Dbname"),
		"dsn":      dsn,
	})

	// fmt.Printf("%s %s %s %d %s",
	// 	viper.GetString("database.username"),
	// 	viper.GetString("database.password"),
	// 	viper.GetString("database.host"),
	// 	viper.GetInt("database.port"),
	// 	viper.GetString("database.Dbname"))

}
