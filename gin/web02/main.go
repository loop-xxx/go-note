package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//post参数
func main() {
	enginp := gin.Default()
	//获取表单数据
	enginp.POST("/login", func(cp *gin.Context) {
		var format string

		uname := cp.PostForm("username")
		passwd := cp.PostForm("password")
		hobby := cp.PostFormArray("hobby")

		if _, exist := cp.GetPostForm("friend"); !exist {
			format = " uname:%s, passwd:%s\n hobby:%v\n default_friend:%s\n"
		} else {
			format = " uname:%s, passwd:%s\n hobby:%v\n friend:%s\n"
		}
		friend := cp.DefaultPostForm("friend", "loop")

		cp.String(http.StatusOK, format, uname, passwd, hobby, friend)
	})
	//表单上传单个文件
	enginp.POST("/upload", func(cp *gin.Context) {
		if fhp, err := cp.FormFile("img"); err == nil{
			if err := cp.SaveUploadedFile(fhp, fmt.Sprintf("up_%s", fhp.Filename)); err == nil{
				cp.String(http.StatusOK, "%s upload success!", fhp.Filename)
			}else{
				cp.String(http.StatusOK, "%v", err)
			}
		}else{
			cp.String(http.StatusOK, "%v", err)
		}
	})
	//表单处理multiPartData
	enginp.POST("/multiPart", func(cp *gin.Context){
		if formp, err := cp.MultipartForm() ; err == nil{
			for _, fhp := range formp.File["data"]{
				if err := cp.SaveUploadedFile(fhp, fmt.Sprintf("up_%s", fhp.Filename)); err == nil{
					fmt.Printf("%s upload success!\n", fhp.Filename)
				}else{
					fmt.Println(err)
				}
			}

			fmt.Println(cp.PostForm("username"))
			for _, item := range formp.Value["username"]{
				fmt.Println(item)
			}
			fmt.Println(cp.PostFormArray("hobby"))
			for _, item := range formp.Value["hobby"]{
				fmt.Println(item)
			}

		}else{
			cp.String(http.StatusOK, "%v", err)
		}
	})
	enginp.Run(":2333")
}
