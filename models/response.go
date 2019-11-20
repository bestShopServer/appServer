package models

//响应信息
type Response struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Total  int    `json:"total"` //查询结果记录数
	Notice string `json:"notice"`
}

//auth.code2Session 腾讯响应信息
type TencentWeChatAuthResp struct {
	OpenId     string `json:"openid"`      //用户唯一标识
	SessionKey string `json:"session_key"` //会话密钥
	UnionId    string `json:"unionid"`     //用户在开放平台的唯一标识符，在满足 UnionID 下发条件的情况下会返回，详见 UnionID 机制说明。
	ErrCode    int    `json:"errcode"`     //错误码
	ErrMsg     string `json:"errmsg"`      //错误信息
}

//type Code2SessionParam struct {
//	OpenId     string `json:"openid"`      //用户唯一标识
//	SessionKey string `json:"session_key"` //会话密钥
//	Token      string `json:"token"`
//}

//auth.code2Session响应信息
type Code2SessionResp struct {
	Response
	Data TencentWeChatAuthResp `json:"data"`
}

//商品列表响应信息
type GoodsListResp struct {
	Response
	Data []TbGoodsBase `json:"data"`
}
