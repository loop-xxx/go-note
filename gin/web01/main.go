package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//uri参数和get参数
func main() {
	// 默认使用了Loger, Recovery两个中间件
	enginp := gin.Default()
	//不使用中间件
	//enginp := gin.New()
	enginp.GET("/", func(cp *gin.Context) {
		cp.String(http.StatusOK, "<h3>hello, web!</h3>")
	})

	//路由匹配规则
	enginp.GET("/rotuer/:user/*tail", func(cp *gin.Context) {
		name := cp.Param("user") //匹配单个
		tail := cp.Param("tail") //匹配剩余
		cp.String(http.StatusOK, "name:%s tail:%s", name, tail)
	})

	//获取参数(?key=value)
	enginp.GET("/query", func(cp *gin.Context) {
		uname := cp.Query("username")
		duname := cp.DefaultQuery("username", "loop")
		cp.String(http.StatusOK, "query:%s\ndefault query:%s\n", uname, duname)
	})

	enginp.Run("0.0.0.0:2333")
}
