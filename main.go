package main

import (
	"fmt"
	"net/http"
	"shop/models"
	"shop/pkg/logging"
	"shop/pkg/setting"
	"shop/routers"
)

func init() {
	//加载配置参数
	setting.Setup()
	//加载数据库
	models.Setup()
}

func main() {

	//清理资源
	defer func() {
		models.CloseDB()
	}()

	routersInit := routers.InitRouter()

	endPoint := fmt.Sprintf("%s:%d",
		"localhost", setting.ServerCfg.HttpPort)

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    setting.ServerCfg.ReadTimeout,
		WriteTimeout:   setting.ServerCfg.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf("[info] start http server listening %s", endPoint)
	logging.Info("服务正常启动 listening:", endPoint)

	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("启动服务ERROR[%v]", err.Error())
	}
}
