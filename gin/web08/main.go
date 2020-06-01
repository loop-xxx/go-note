package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginFilterFunc(cp *gin.Context){
	if _, err:=cp.Cookie("user"); err != nil{
		cp.Redirect(http.StatusTemporaryRedirect,"http://localhost:63342/loop-xxx/web08/template/index.html")
		//不再调用后续的函数处理
		cp.Abort()
	}

}

func main(){
	enginep := gin.Default()

	enginep.POST("/login", func(cp *gin.Context){
		cp.SetCookie("user", cp.DefaultPostForm("username", "loop"),
			60, "/home", "localhost",
			false, //是否只支持https
			 true)//是否不支持js访问
		cp.XML(http.StatusOK, gin.H{"status":"success", "home":"http://localhost:2333/home/index"})
	})

	home := enginep.Group("/home")
	home.Use(LoginFilterFunc)
	home.GET("/:name", func(cp *gin.Context){
		username, _:= cp.Cookie("user")
		cp.JSON(http.StatusOK, gin.H{
			"user": username,
			fmt.Sprintf("%s_ok", cp.Param("name")):"true",
		})
	})
	home.GET("/", func(cp *gin.Context){
		username, _:= cp.Cookie("user")
		cp.JSON(http.StatusOK, gin.H{
			"user": username,
		})
	})


	enginep.Run(":2333")
}
