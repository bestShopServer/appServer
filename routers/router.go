package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/pkg/logging"
	"shop/pkg/setting"
	"shop/services"
)

func MiddleWare() gin.HandlerFunc {

	return func(c *gin.Context) {
		//写日志
		//_ , _=fmt.Fprintln(gin.DefaultWriter, fmt.Sprintf("%s |%s |%s |%s |%s |%v ",
		//	time.Now().Format(time.RFC3339Nano),
		//	c.ClientIP(),
		//	c.Request.Proto,
		//	c.Request.Method,
		//	c.Request.UserAgent(),
		//	c.Request.URL,
		//))
		logging.Info("\n", c.Request.URL, " IP:", c.ClientIP(), "|", c.Request.Proto,
			"|", c.Request.Method, "|", c.Request.UserAgent())

		Authorization := c.GetHeader("Authorization")
		logging.Info("Header Authorization:", Authorization)

		if c.Request.Method == "GET" {
			logging.Info("GET:", c.Request.URL)
		} else if c.Request.Method == "POST" {
			//打印body日志
			//body,_ := ioutil.ReadAll( c.Request.Body )
			//utils.DebugLog("POST[%v]", string( body ) )
		}
		c.Next()

		logging.Info(c.Request.URL.Path, " Handle END")
	}
}

func InitRouter() *gin.Engine {

	gin.SetMode(setting.ServerCfg.RunMode)

	r := gin.New()
	r.Use(gin.Logger(), MiddleWare(), gin.Recovery())

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
