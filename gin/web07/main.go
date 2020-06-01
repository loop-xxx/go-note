package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/**
 middleware
 Next(), 是否调用, 决定中间件执行策略是嵌套还是顺序
 */

func Filter() func (*gin.Context){
	// init Filter codes
	return func (cp *gin.Context){
		start := time.Now()
		fmt.Printf("Filter:%s\n", cp.Request.URL)

		//cp.Next() //决定执行的策略是嵌套, 还是顺序

		fmt.Printf("Fileter:%v\n", time.Since(start))
	}
}

func FilterFunc( cp *gin.Context){
	start := time.Now()
	fmt.Printf("FilterFunc:%s\n", cp.Request.URL)

	cp.Next()

	fmt.Printf("FilterFunc:%v\n", time.Since(start))
}

func GroupFilter() func (*gin.Context){
	// init Filter codes
	return func (cp *gin.Context){
		start := time.Now()
		fmt.Printf("Group:%s\n", cp.Request.URL)

		cp.Next()

		fmt.Printf("Group:%v\n", time.Since(start))
	}
}

func GroupFilterFunc(cp *gin.Context){
	start := time.Now()
	fmt.Printf("groupFunc:%s\n", cp.Request.URL)

	//cp.Next()

	fmt.Printf("groupFunc:%v\n", time.Since(start))
}

func main(){
	enginep := gin.Default()

	//全局middleware

	enginep.Use(Filter())
	enginep.Use(FilterFunc)

	enginep.GET("/index", func( cp *gin.Context){
		fmt.Println("running")
		cp.JSON(http.StatusOK, gin.H{"index_ok": "true"})
	})

	//路由组middleware
	play  := enginep.Group("/play", GroupFilter())
	play.Use(GroupFilterFunc)
	play.GET("/:name", func(cp *gin.Context){
		fmt.Println("running")
		cp.JSON(http.StatusOK,gin.H{cp.Param("name"):"true"})
	})


	//单路由middleware
	enginep.GET("/filter", func(cp *gin.Context){
		start := time.Now()
		fmt.Printf("single1:%s\n", cp.Request.URL)

		//cp.Next()

		fmt.Printf("single1:%v\n", time.Since(start))

	}, func (cp *gin.Context){
		start := time.Now()
		fmt.Printf("single2:%s\n", cp.Request.URL)

		cp.Next()

		fmt.Printf("single2:%v\n", time.Since(start))

	}, func(cp *gin.Context){
		fmt.Println("running")
		cp.JSON(http.StatusOK, gin.H{"single_ok":"true"})
	})

	enginep.Run("localhost:2333")
}
