package main

import (
	"fmt"
	"github.com/loop-xxx/go-note/gin/web10/control"
	"github.com/loop-xxx/go-note/gin/web10/modle"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := modle.InitDB(); err == nil {
		enginep := gin.Default()
		enginep.LoadHTMLGlob("./view/template/*")
		enginep.Static("/static", "./view/static")
		enginep.GET("/", control.RootFunc)
		enginep.GET("/list", control.ListFunc)
		enginep.GET("/delete/:uuid", control.DeleteFunc)
		enginep.POST("/add", control.AddFunc)
		_ = enginep.Run(":2333")

		modle.CloseDB()
	} else {
		fmt.Println(err)
	}
}
