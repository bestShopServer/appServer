module shop

go 1.12

require (
	github.com/gin-gonic/gin v1.4.0
	github.com/go-ini/ini v1.50.0
	github.com/jinzhu/gorm v1.9.11
	github.com/kr/pretty v0.1.0 // indirect
	github.com/unknwon/com v1.0.1
	google.golang.org/appengine v1.6.5 // indirect
	gopkg.in/ini.v1 v1.51.0 // indirect
)

replace (
	google.golang.org/appengine => github.com/golang/appengine v1.6.5
	shop/conf => ./pkg/conf
	shop/middleware => ./middleware
	shop/models => ./models
	shop/pkg/logging => ./pkg/logging
	shop/pkg/setting => ./pkg/setting
	shop/routers => ./routers
)
