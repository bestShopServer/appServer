package models

import "shop/pkg/setting"

//通过页码计算limit起始值和步长
func PageGetLimit(page int, ostep int) (offset int, step int) {
	if ostep == 0 {
		step = setting.Dbcfg.Step
	} else {
		step = ostep
	}
	return page * step, step
}
