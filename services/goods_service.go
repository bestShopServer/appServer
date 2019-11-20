package services

import (
	"github.com/gin-gonic/gin"
	"shop/middleware/app"
	"shop/models"
	"shop/pkg/e"
	"shop/pkg/logging"
)

//获取商品列表
func PostGoodsList(c *gin.Context) {
	appG := app.Gin{C: c}
	logging.Info("PostGoodsList ...")

	var jparm models.GoodsListReq
	err := c.ShouldBindJSON(&jparm)
	logging.Info("请求参数:", jparm)

	// Optional Fields List
	//var optionalFields []string
	optionalFields := []string{}

	// Check Param
	if !CheckParams(jparm, "PostGoodsList", err, optionalFields) {
		appG.ResponseJsonError(e.INVALID_PARAMS)
		return
	}

	res, total, err := models.QueryGoodsList(jparm)
	if err != nil {
		logging.Error("ERROR[%v]", err.Error())
		appG.ResponseJsonError(e.ERROR)
		return
	}
	logging.Info("查询商品列表成功!")

	//准备响应信息
	resp := models.GoodsListResp{}
	resp.Code = e.SUCCESS
	resp.Total = total
	resp.Data = res
	appG.ResponseJsonMessage(resp)
}
