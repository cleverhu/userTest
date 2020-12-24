package handlers

import (
	"github.com/gin-gonic/gin"
	"time"
	"userTest/src/data/Getter"
	"userTest/src/data/Setter"
	"userTest/src/models/UserModel"
	"userTest/src/result"
)

func GetUserList(ctx *gin.Context) {
	R(ctx)("10000", "query users success", Getter.UserGetter.GetUserList())(OK)
}

func AddUser(ctx *gin.Context) {
	u := UserModel.New().Mutate(UserModel.WithUpdateTime(time.Now()))
	result.Result(ctx.ShouldBindJSON(u)).Unwrap()
	R(ctx)("10001", "add user success", Setter.UserSetter.AddUser(u).Unwrap())(OK)
}

func GetLogList(ctx *gin.Context) {
	R(ctx)("10000", "query logs success", Getter.LogGetter.GetLogList())(OK)
}
