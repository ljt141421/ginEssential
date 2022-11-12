package common

import (
	"fmt"
	"ginEssential/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initDB() *gorm.DB {

	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	dsn := fmt.Sprintf("%s:%st@tcp(%s:%s)/%sl?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, database)
	//dsn := "root:root@tcp(101.43.169.25:3306)/personal?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("fail to connect database, err: " + err.Error())
	}
	db.AutoMigrate(&model.User{})

	return db

}

func GetDB() *gorm.DB {
	DB = initDB()
	return DB
}
