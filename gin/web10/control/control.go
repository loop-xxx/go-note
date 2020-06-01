package control

import (
	"github.com/loop-xxx/go-note/gin/web10/modle"
	"github.com/loop-xxx/go-note/gin/web10/modle/bean"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func ListFunc(cp *gin.Context) {
	urlList := modle.GetUrlList()
	cp.HTML(http.StatusOK, "list.html", gin.H{"urlList": urlList})
}

func AddFunc(cp *gin.Context) {
	//StatusTemporaryRedirect, StatusPermanentRedirect 保持原来的请求数据重新发送请求到重定向地址
	//StatusMovedPermanently 重新发送一个新GET请求到重定向地址
	var url bean.Url
	if err := cp.Bind(&url); err == nil {

		if t, err := time.Parse("2006/01/02", url.Date); err == nil {
			url.Date = t.Format("2006-01-02")
		} else {
			log.Println(err)
		}
		url.UUID = uuid.NewV4().String()
		modle.AddUrl(&url)

		cp.Redirect(http.StatusMovedPermanently, "/list")
	} else {
		log.Println(err)
		cp.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "info": "http.StatusBadRequest"})
	}
}

func RootFunc(cp *gin.Context) {
	cp.Redirect(http.StatusPermanentRedirect, "/list")
}

func DeleteFunc(cp *gin.Context) {
	if urlUUID := cp.Param("uuid"); urlUUID != "" {
		modle.DelByUUID(urlUUID)
	}
	cp.Redirect(http.StatusMovedPermanently, "/list")
}
