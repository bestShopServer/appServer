package app

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"shop/pkg/logging"
	"shop/pkg/setting"
)

func HttpRequest(url string) (body []byte, err error) {
	logging.Info("Get URL[%s]", url)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		logging.Error("Get URL Error [%v]", err.Error())
		return body, err
	}
	body, err = ioutil.ReadAll(resp.Body)
	logging.Info("body[%s]", string(body))
	return body, err
}

//调用小程序登录接口
func WeChatLogin(code string) ([]byte, error) {
	//https://api.weixin.qq.com/sns/jscode2session?appid=APPID
	// &secret=SECRET&js_code=JSCODE&grant_type=authorization_code
	base_url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%v"+
		"&secret=%v&js_code=%v&grant_type=authorization_code",
		setting.WechatCfg.AppId, setting.WechatCfg.Secret, code)
	body, err := HttpRequest(base_url)
	return body, err
}
