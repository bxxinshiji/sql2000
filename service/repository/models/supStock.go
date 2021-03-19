package models

import (
	"github.com/MXi4oyu/Utils/cnencoder/gbk"
	"github.com/bxxinshiji/sql2000/util"
)

// SupStock 部门库存
type SupStock struct {
	SupCode string `xorm:"comment('部门编码') 'SupCode'"`
	PluCode string `xorm:"comment('商品编码') 'PluCode'"`
	Number string `xorm:"comment('库存数量') 'Number'"`
	Name string `xorm:"comment('部门名称') 'Name'"`
}

// Hander 数据预处理
func (b *SupStock) Hander() {
	b.SupCode = util.TrimSpace(b.SupCode)
	b.PluCode = util.TrimSpace(b.PluCode)
	b.Number = util.TrimSpace(b.Number)
	b.Name = util.TrimSpace(gbk.Decode(b.Name))
}
