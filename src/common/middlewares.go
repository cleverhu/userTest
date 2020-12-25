package common

import (
	"github.com/gin-gonic/gin"

)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				ctx.JSON(400, gin.H{"message": err})
			}
		}()

		ctx.Next()
	}
}


func AuthMiddleWare() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		defer func() {
			//ctx.GetHeader()
			//ParseToken()
			if err := recover(); err != nil {
				ctx.JSON(400, gin.H{"message": err})
			}
		}()

		ctx.Next()
	}
}