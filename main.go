package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(11);not null;unique"`
	Password  string `gorm:"size:255;not null"`
}

func main() {

	db := InitDB()
	r := gin.Default()
	r.POST("/api/auth/register", func(ctx *gin.Context) {

		name := ctx.PostForm("name")
		telephone := ctx.PostForm("telephone")
		password := ctx.PostForm("password")

		if len(telephone) != 11 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "message": "手机号必须是11位"})
			return
		}

		if len(password) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "message": "密码不能少于六位"})
			return
		}

		//如果没有传name，则生成一个随机的十位字符串name
		if len(name) == 0 {
			name = RandomString(10)
		}

		//检查手机号是否重复存在
		if isTelephoneExist(db, telephone) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "message": "手机号已经存在"})
			return
		}

		log.Println(name, password, telephone)

		newUser := User{
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}
		db.Create(&newUser)

		ctx.JSON(200, gin.H{"code": 200, "message": "成功"})
	})
	r.Run() // listen and serve on 0.0.0.0:8080

}

func RandomString(n int) string {
	letters := []byte("agasngsjsgehgsesfsgsegesasgaeasglghjhns")
	result := make([]byte, n)

	//给随机函数一个种子
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func isTelephoneExist(db *gorm.DB, phone string) bool {
	var user User
	db.Where("telephone = ?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func InitDB() *gorm.DB {
	//driverName := "mysql"
	//host := "101.43.169.25"
	//port := "3306"
	//database := "personal"
	//username := "root"
	//password := "root"
	//charset := "utf8"
	//args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
	//	username,
	//	password,
	//	host,
	//	port,
	//	database,
	//	charset)

	dsn := "root:root@tcp(101.43.169.25:3306)/personal?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("fail to connect database, err: " + err.Error())
	}
	db.AutoMigrate(&User{})
	return db

}
