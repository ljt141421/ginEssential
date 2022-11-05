package common

import (
	"ginEssential/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initDB() *gorm.DB {

	dsn := "root:root@tcp(101.43.169.25:3306)/personal?charset=utf8mb4&parseTime=True&loc=Local"
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
