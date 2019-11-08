

package main

import (
    "fmt"
	"net/http"
	"shop/pkg/setting"
	"shop/routers"
	"shop/pkg/logging"
)

func main() {
    router := routers.InitRouter()

	logging.Info("server start")
    s := &http.Server{
        Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
        Handler:        router,
        ReadTimeout:    setting.ReadTimeout,
        WriteTimeout:   setting.WriteTimeout,
        MaxHeaderBytes: 1 << 20,
    }

    s.ListenAndServe()
}

