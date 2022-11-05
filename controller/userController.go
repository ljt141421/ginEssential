package controller

import (
	"ginEssential/common"
	"ginEssential/model"
	"ginEssential/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var db = common.GetDB()

// Register 用户注册
func Register(ctx *gin.Context) {

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

	//加密密码
	hasePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "手机号已经存在"})
		return
	}

	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasePassword),
	}
	db.Create(&newUser)

	ctx.JSON(200, gin.H{"code": 200, "message": "成功"})
}

// Login 用户登录
func Login(ctx *gin.Context) {
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

	//判断手机号是否存在
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "message": "手机号不存在"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 500, "message": "密码错误"})
		return
	}

	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 500, "message": "系统异常"})
		log.Printf("token generate error: %v", err)
		return
	}
	ctx.JSON(200, gin.H{"code": 200, "data": gin.H{"token": token}, "message": "登录成功"})

}

// Info 获取用户信息
func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"user": user},
	})
}

func isTelephoneExist(db *gorm.DB, phone string) bool {
	var user model.User
	db.Where("telephone = ?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
