package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type person struct{
	Name string `json:"name" xml:"name" yaml:"name"`
	Age uint32 `json:"age" xml:"age" yaml:"age"`
}

//gin response 支持各种数据类型
func main(){
	enginep:=gin.Default()
	enginep.GET("/json", func(cp *gin.Context){
		p:=person{Name:"loop", Age:25}
		cp.JSON(http.StatusOK, p)
	})

	enginep.POST("/json", func(cp *gin.Context){
		cp.JSON(http.StatusOK, gin.H{"name":"loop", "age":25})
	})

	enginep.GET("/xml", func(cp *gin.Context){
		p:=person{Name:"loop", Age:25}
		cp.XML(http.StatusOK, p)
	})

	enginep.POST("/xml", func(cp *gin.Context){
		cp.XML(http.StatusOK, gin.H{"name":"loop", "age":25})
	})

	enginep.GET("/yaml", func(cp *gin.Context){
		p:=person{Name:"loop", Age:25}
		cp.YAML(http.StatusOK, p)
	})

	enginep.POST("/yaml", func(cp *gin.Context){
		cp.YAML(http.StatusOK, gin.H{"name":"loop", "age":25})
	})


	enginep.Run(":2333")
}
