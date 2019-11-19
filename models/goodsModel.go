package models

import "shop/pkg/logging"

//参阅文档 https://jasperxu.github.io/gorm-zh/crud.html#q
func QueryGoodsList(param GoodsListReq) (goods []TbGoodsBase, total int, err error) {
	logging.Debug("QueryGoodsList ...")
	offset, step := PageGetLimit(param.Pages, param.Step)

	res := db.Table(TbGoodsBase{}.TableName()).
		Where(" goods_type_id=?", param.GoodsTypeId).
		Order(" d.goods_id desc ")
	res.Count(&total)
	logging.Info("SQL[%v]结果数[%v]", res.QueryExpr(), total)

	if total == 0 { //数据为0时，直接return
		return goods, total, err
	}
	//处理分页
	res = res.Limit(step).Offset(offset).Find(&goods)
	logging.Info("SQL[%v]", res.QueryExpr())

	if res.Error != nil {
		logging.Error("select goods base error[%v]", res.Error.Error())
		return goods, total, res.Error
	}
	logging.Info("商品数量[%v]", total)
	return goods, total, err
}
