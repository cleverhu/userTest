package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"userTest/src/common"
	"userTest/src/data/Getter"
	"userTest/src/data/Setter"
	"userTest/src/models/UserModel"
	"userTest/src/result"
)

func GetUserList(ctx *gin.Context) {
	R(ctx)("10000", "query users success", Getter.UserGetter.GetUserList())(OK)
}

func AddUser(ctx *gin.Context) {
	u := UserModel.New().Mutate(UserModel.WithUpdateTime(UserModel.MyTime(time.Now())))
	result.Result(ctx.ShouldBindJSON(u)).Unwrap()
	R(ctx)("10001", "add user success", Setter.UserSetter.AddUser(u).Unwrap())(OK)
}

func GetLogList(ctx *gin.Context) {
	R(ctx)("10002", "query logs success", Getter.LogGetter.GetLogList())(OK)
}

func Login(ctx *gin.Context) {
	u := UserModel.NewUserLoginInfoImpl()
	result.Result(ctx.ShouldBindJSON(u)).Unwrap()
	R(ctx)("10003", "login success", Getter.UserGetter.Login(u).Unwrap())(OK)
}

func LoginGet(ctx *gin.Context) {
	data := make(map[string]interface{})
	err := fmt.Errorf("")
	token, _ := ctx.Cookie("jwt")
	if token == "" {
		query := ctx.Query("token")
		if data, err = common.ParseToken(query); err == nil {
			ctx.SetCookie("jwt", query, 3600, "/", "/", false, true)
			ctx.HTML(200, "index.html", gin.H{"data": data})
		} else {
			ctx.Redirect(http.StatusFound, "http://auth.deeplythink.com/?redirect_url=http://www.deeplythink.com")
		}
	} else {
		if data, err = common.ParseToken(token); err == nil {
			ctx.HTML(200, "index.html", gin.H{"data": data})
		} else {
			ctx.Redirect(http.StatusFound, "http://auth.deeplythink.com/?redirect_url=http://www.deeplythink.com")
		}
	}

	//u := UserModel.NewUserLoginInfoImpl()
	//result.Result(ctx.ShouldBindJSON(u)).Unwrap()
	//R(ctx)("10003", "login success", Getter.UserGetter.Login(u).Unwrap())(OK)

}
