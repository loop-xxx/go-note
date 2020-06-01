package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type person struct{
	Name string `json:"name"`
	Age uint32 `json:"age"`
}

//路由组
func main(){
	p := person{"loop", 25}
	enginp := gin.Default()
	play:= enginp.Group("/play")
	play.GET("/game/:name", func(cp *gin.Context){
		fmt.Println(cp.Param("name"))
		fmt.Println(cp.Query("name"))
		//cp.JSON(http.StatusOK, gin.H{"person": p})
		cp.JSON(http.StatusOK, map[string] person{"person":p})
	})
	play.POST("/animation/:name", func(cp *gin.Context){
		fmt.Println(cp.Param("name"))
		fmt.Println(cp.PostForm("name"))
		cp.JSON(http.StatusOK, p)
	})

	dance:=enginp.Group("/dance")
	dance.GET("/v1", func(cp *gin.Context){
		cp.String(http.StatusOK, "dance_v1")
	})
	dance.GET("/v2", func(cp *gin.Context){
		cp.String(http.StatusOK, "dance_v2")
	})

	enginp.Run(":2333")
}
