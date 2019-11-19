package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/pkg/setting"
	"shop/services"
)

func InitRouter() *gin.Engine {

	gin.SetMode(setting.ServerCfg.RunMode)

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	apiv1 := r.Group("/api/v1")
	{
		//test
		apiv1.GET("/test", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})
		apiv1.POST("/code2session", services.PostCode2Session)
		apiv1.POST("/goods/list", services.PostGoodsList)
	}
	return r
}
