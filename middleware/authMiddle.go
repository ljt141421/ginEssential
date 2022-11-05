package middleware

import (
	"ginEssential/common"
	"ginEssential/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// AuthMiddleAware 用户认证中间件
func AuthMiddleAware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		//token无效
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		//解析token失败
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}
		userId := claims.UserId
		db := common.GetDB()
		var user model.User
		db.First(&user, userId)
		//验证用户是否存在
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}
		//用户存在 将user信息写入上下文
		ctx.Set("user", user)
		ctx.Next()

	}
}
