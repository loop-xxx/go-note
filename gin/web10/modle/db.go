package modle

import (
	"database/sql"
	"github.com/loop-xxx/go-note/gin/web10/modle/bean"
	"log"
)
import _ "github.com/go-sql-driver/mysql"

var DBClient *sql.DB

func InitDB()(err error){
	dataSourceName := "root:toor@tcp(127.0.0.1:3306)/web01"
	if DBClient, err = sql.Open("mysql", dataSourceName); err == nil{
		if err = DBClient.Ping(); err != nil{
			_ = DBClient.Close()
		}
	}
	return
}

func CloseDB(){
	_ = DBClient.Close()
	return
}

func AddUrl(urlp *bean.Url){
	sqlStr := "insert into urls(u_uuid, u_title, u_url, u_date, u_desc) values(?, ?, ?, ?, ?)"
	if _, err := DBClient.Exec(sqlStr, urlp.UUID, urlp.Title, urlp.Url, urlp.Date, urlp.Desc); err != nil{
		panic(err)
	}
}


func DelByUUID(uuid string){
	sqlStr := "delete from urls where u_uuid = ?"
	if _, err := DBClient.Exec(sqlStr, uuid); err != nil{
		panic(err)
	}
}
func GetUrlList()(urlList []bean.Url){
	urlList = make([]bean.Url,0, 0x10)
	var url bean.Url

	sqlStr := "select u_uuid, u_title, u_url, u_date, u_desc from urls;"

	if rows, err := DBClient.Query(sqlStr); err == nil{
		defer rows.Close()
		for rows.Next(){
			if err := rows.Scan(&url.UUID, &url.Title, &url.Url, &url.Date, &url.Desc); err == nil{
				urlList = append(urlList, url)
			}else {
				log.Println(err)
			}
		}
	}else{
		panic(err)
	}
	return
}