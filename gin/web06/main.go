package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func handleFunc(cp *gin.Context){
	time.Sleep(time.Second*4)
	fmt.Println(cp.Request.URL)
}

func main(){
	//html渲染
	enginep := gin.Default()
	//加载模板, 使用通配符
	//enginep.LoadHTMLGlob("template/*")
	//加载模板, 使用文件列表
	enginep.LoadHTMLFiles("./template/index.html")
	enginep.GET("/index/:name", func(cp *gin.Context){
		cp.HTML(http.StatusOK, "index.html", map[string]string{"xxx":cp.Param("name")})
	})


	//请求重定向
	enginep.GET("/redirect", func (cp *gin.Context){
		cp.Redirect(http.StatusTemporaryRedirect, "https://www.baidu.com")
	})


	//请求异步处理
	enginep.GET("/async", func(cp *gin.Context){
		go handleFunc(cp.Copy()) //异步不能使用原先的Context对象, 使用Copy()函数拷贝出一个副本
		cp.JSON(http.StatusOK, gin.H{"async_ok":"true"})
	})
	//请求同步处理
	enginep.GET("/sync", func(cp *gin.Context){
		handleFunc(cp)
		cp.JSON(http.StatusOK, gin.H{"sync_ok":"true"})
	})

	enginep.Run(":2333")
}
