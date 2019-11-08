module shop

go 1.13

require (
	github.com/astaxie/beego v1.12.0
	github.com/gin-gonic/gin v1.4.0
	github.com/go-ini/ini v1.50.0
	github.com/jinzhu/gorm v1.9.11
	github.com/unknwon/com v1.0.1
	google.golang.org/appengine v1.6.5 // indirect
)

replace (
	shop/conf => ./pkg/conf
	shop/middleware => ./middleware
	shop/models => ./models
	shop/pkg/logging => ./pkg/logging
	shop/pkg/setting => ./pkg/setting
	shop/routers => ./routers
)
