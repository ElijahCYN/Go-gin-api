module github.com/ElijahCYN/Go-gin-api

go 1.16

replace github.com/ElijahCYN/Go-gin-api/pkg/setting => /User/Documents/Golang/src/SideProject/Go-gin-api/pkg/setting

require (
	github.com/astaxie/beego v1.12.3
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.62.0
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/unknwon/com v1.0.1
	gopkg.in/ini.v1 v1.62.0 // indirect
)
