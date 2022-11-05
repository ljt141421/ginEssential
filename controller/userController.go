package controller

import (
	"ginEssential/common"
	"ginEssential/model"
	"ginEssential/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Register(ctx *gin.Context) {

	db := common.GetDB()

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
		name = utils.RandomString(10)
	}

	//检查手机号是否重复存在
	if isTelephoneExist(db, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "message": "手机号已经存在"})
		return
	}

	log.Println(name, password, telephone)

	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	db.Create(&newUser)

	ctx.JSON(200, gin.H{"code": 200, "message": "成功"})
}

func isTelephoneExist(db *gorm.DB, phone string) bool {
	var user model.User
	db.Where("telephone = ?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
