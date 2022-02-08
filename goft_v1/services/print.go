package services

import "github.com/gin-gonic/gin"

func Print(ctx *gin.Context) {
	ctx.String(200, "%s 成功打印", "hacker")
}
