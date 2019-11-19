package services

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"shop/middleware/app"
	"shop/models"
	"shop/pkg/e"
	"shop/pkg/logging"
)

//微信Code2Session获取openid
func PostCode2Session(c *gin.Context) {
	appG := app.Gin{C: c}
	logging.Info("PostCode2Session ...")

	var jparm models.Code2SessionReq
	var tparm models.TencentWeChatAuthResp
	err := c.ShouldBindJSON(&jparm)
	//err := c.ShouldBindQuery(&jparm)
	if err != nil {
		appG.ResponseJsonError(e.INVALID_PARAMS)
		return
	}
	logging.Info("请求参数:", jparm)

	// Optional Fields List
	//var optionalFields []string
	optionalFields := []string{"Location"}

	// Check Param
	if !CheckParams(jparm, "Code2Session", err, optionalFields) {
		appG.ResponseJsonError(e.INVALID_PARAMS)
		return
	}

	//请求腾讯登录接口
	res, err := app.WeChatLogin(jparm.Code)
	if err != nil {
		appG.ResponseJsonError(e.INVALID_PARAMS)
		return
	}
	//解析腾讯接口返回的json报文
	err = json.Unmarshal(res, &tparm)
	if err != nil {
		appG.ResponseJsonError(e.INVALID_PARAMS)
		return
	}

	if tparm.ErrCode != 0 {
		appG.ResponseJsonError(e.INVALID_PARAMS)
		return
	}

	//准备响应信息
	rtn := models.Code2SessionResp{}
	rtn.Data = tparm
	appG.ResponseJsonMessage(rtn)
}
