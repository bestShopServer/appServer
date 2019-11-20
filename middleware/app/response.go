package app

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/pkg/e"
	"shop/pkg/logging"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}

// Response setting gin.JSON
func (g *Gin) ResponseJsonSuccess() {
	g.C.JSON(http.StatusOK, Response{
		Code: e.SUCCESS,
		Msg:  e.GetMsg(e.SUCCESS),
	})
	logging.Info("%v", http.StatusOK)
	return
}

// Response setting gin.JSON
func (g *Gin) ResponseJsonError(code int) {
	g.C.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  e.GetMsg(code),
	})
	logging.Info("%v", http.StatusOK)
	return
}

// Response setting gin.JSON
func (g *Gin) ResponseJsonMessage(obj interface{}) {
	g.C.JSON(http.StatusOK, obj)
	logging.Info("%+v", obj)
	return
}

// Response setting gin.XML
func (g *Gin) ResponseXmlSuccess() {
	g.C.JSON(http.StatusOK, Response{
		Code: e.SUCCESS,
		Msg:  e.GetMsg(e.SUCCESS),
	})
	logging.Info("%v", http.StatusOK)
	return
}

// Response setting gin.XML
func (g *Gin) ResponseXmlMessage(obj interface{}) {
	by, _ := xml.Marshal(obj)
	logging.Info("响应信息[%v]", string(by))
	g.C.XML(http.StatusOK, obj)
	return
}
