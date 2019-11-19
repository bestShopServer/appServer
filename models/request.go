package models

type CommonReq struct {
	OpenId string `json:"open_id"`
}

type PagingReq struct {
	Pages int `json:"pages"`
	Step  int `json:"step"`
}

//用户auth.code2Session
type Code2SessionReq struct {
	Code string `json:"code"  form:"code"`
}

//用户auth.code2Session
type GoodsListReq struct {
	CommonReq
	PagingReq
	GoodsTypeId int `json:"goods_type_id"`
}
