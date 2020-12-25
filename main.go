package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"userTest/src/common"
	_ "userTest/src/dbs"
	"userTest/src/handlers"
	_ "userTest/src/myValidators"
)

func main() {

	r := gin.New()

	r.Use(common.ErrorHandler())
	r.LoadHTMLGlob("src/html/*")

	r.GET("/users", handlers.GetUserList)
	r.GET("/logs", handlers.GetLogList)
	r.POST("/users", handlers.AddUser)

	r.POST("/login", handlers.Login)

	r.GET("/login", handlers.LoginGet)
	r.GET("/", handlers.LoginGet)

	log.Fatal(r.Run(":80"))
}
