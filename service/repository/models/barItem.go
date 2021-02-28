package models

import (
	"github.com/MXi4oyu/Utils/cnencoder/gbk"
	"github.com/lecex/sql2000/util"
)

// BarItem 商品
type BarItem struct {
	BarCode string `xorm:"comment('条形码') 'BarCode'"`
	PluCode string `xorm:"comment('商品编码') 'PluCode'"`
	DepCode string `xorm:"comment('部门') 'DepCode'"`
	PluName string `xorm:"comment('商品名称') 'PluName'"`
	Spec    string `xorm:"comment('规格') 'Spec'"`
}

// Hander 数据预处理
func (b *BarItem) Hander() {
	b.PluCode = util.TrimSpace(b.PluCode)
	b.BarCode = util.TrimSpace(b.BarCode)
	b.DepCode = util.TrimSpace(b.DepCode)
	b.PluName = util.TrimSpace(gbk.Decode(b.PluName))
	b.Spec = util.TrimSpace(gbk.Decode(b.Spec))
}
