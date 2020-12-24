package handlers

import (
	"github.com/gin-gonic/gin"
	"sync"
)

type JSONResult struct {
	Code    string
	Message string
	Result  interface{}
}

var ResultPool *sync.Pool

func init() {
	ResultPool = &sync.Pool{New: func() interface{} {
		return gin.H{"code": "", "message": "", "result": nil}
	}}
}

//R(ctx *gin.Context)(code,message string,result interface)(ctx *gin.context,{}interface)
type Output func(ctx *gin.Context, data interface{})
type ResultFunc func(code, message string, result interface{}) func(Output) //中间的函数

//
func R(ctx *gin.Context) ResultFunc {
	return func(code, message string, result interface{}) func(Output) {
		data := ResultPool.Get().(gin.H)
		ResultPool.Put(data)
		data["code"] = code
		data["message"] = message
		data["result"] = result
		return func(output Output) {
			output(ctx, data)
		}
	}
}

func OK(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, data)
}

func Error(ctx *gin.Context, data interface{}) {
	ctx.JSON(400, data)
}
