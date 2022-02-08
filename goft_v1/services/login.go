package services

import "github.com/gin-gonic/gin"

func Login(ctx *gin.Context) {
	//nick := ctx.DefaultPostForm("nick", "anonymous")
	ctx.JSON(200, gin.H{
		"status":  "posted",
		"message": "登陆成功",
		//"nick":    nick,
	})
}
