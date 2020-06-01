package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)


//binding tag 必须参数
type person struct{
	Name string `json:"name" form:"name" uri:"name" binding:"required"`
	Age int32 `json:"age" form:"age" uri:"age" binding:"required"`
}

//数据解析和邦定
func main(){
	enginep := gin.Default()

	//根据uri携带的参数解析数据
	enginep.GET("/uri/:name/:age", func(cp *gin.Context){
		var p person
		if err := cp.ShouldBindUri(&p); err == nil{
			fmt.Println(p)
			cp.String(http.StatusOK, "uri_get_success")
		}
	})
	enginep.POST("/uri/:name/:age", func(cp *gin.Context){
		var p person
		if err := cp.ShouldBindUri(&p); err == nil{
			fmt.Println(p)
			cp.String(http.StatusOK, "uri_post_success")
		}
	})

	//根据url 携带的query参数解析数据
	enginep.GET("/query", func(cp *gin.Context){
		var p person
		if err := cp.ShouldBindQuery(&p); err == nil{
			fmt.Println(p)
			cp.String(http.StatusOK, "query_get_success")
		}
	})

	enginep.POST("/query", func(cp *gin.Context){
		var p person
		if err := cp.ShouldBindQuery(&p); err == nil{
		fmt.Println(p)
		cp.String(http.StatusOK, "query_post_success")
		}
	})

	//post请求, 根据body中json格式的数据 解析数据
	enginep.POST("/json", func(cp *gin.Context){
		var p person
		if err := cp.ShouldBindJSON(&p); err == nil{
		fmt.Println(p)
		cp.String(http.StatusOK, "json_success!")
		}
	})

	//post请求, 根据body中的数据(根据请求头Content-Type辨别数据格式 支持json) 解析数据
	enginep.POST("/form", func(cp *gin.Context){
		var p person
		//Bind解析缺少必须参数时, 会设置400错误码, 其他几种解析方式(ShouldBindXXX)只会返回err,不会设置
		if err := cp.Bind(&p); err == nil{
		fmt.Println(p)
		cp.String(http.StatusOK, "form_success!")
		}
	})

	if err := enginep.Run(":2333"); err != nil{
		fmt.Println(err)
	}
}