package models

import (
	"time"
)

type TbUserType struct {
	TypeID   int64     `gorm:"primary_key" json:"type_id"`
	TypeName string    `json:"type_name"`
	Thumb    string    `json:"thumb"`
	InfoTime time.Time `gorm:"default:auto_now_add" json:"info_time"`
}

//表名
func (TbUserType) TableName() string {
	return "user_type"
}

//唯一索引
func (TbUserType) TableUnique() [][]string {
	return [][]string{
		[]string{"TypeId"},
	}
}

type TbGoodsBase struct {
	GoodsID     int64  `gorm:"primary_key" json:"type_id"`
	GoodsName   string `json:"type_name"`
	GoodsTypeId int    `json:"goods_type_id"`
	UserTypeId  int
	Head        string
	Thumb       string `json:"thumb"`
	Price       float64
	Discount    float64
	SalePrice   float64
	Details     string
	Specs       string
	Service     string
	Remark      string
	InfoTime    time.Time `gorm:"default:auto_now_add" json:"info_time"`
}

//表名
func (TbGoodsBase) TableName() string {
	return "goods_base"
}
