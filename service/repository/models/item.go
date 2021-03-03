package models

import (
	"github.com/MXi4oyu/Utils/cnencoder/gbk"
	"github.com/bxxinshiji/sql2000/util"
)

// Item 商品
type Item struct {
	PluCode     string `xorm:"comment('商品ID') 'PluCode'"`
	PluName     string `xorm:"comment('商品名称') 'PluName'"`
	PluAbbrName string `xorm:"comment('商品名称') 'PluAbbrName'"`
	BarCode     string `xorm:"comment('条形码') 'BarCode'"`
	Spec        string `xorm:"comment('商品规格') 'Spec'"`
	SPrice      string `xorm:"comment('商品规格') 'SPrice'"`
	HJPrice     string `xorm:"comment('商品规格') 'HJPrice'"`
	Unit        string `xorm:"comment('计量单位') 'Unit'"`
	DepCode     string `xorm:"comment('部门编码') 'DepCode'"`
	ClsCode     string `xorm:"comment('部门编码') 'ClsCode'"`
	SupCode     string `xorm:"comment('未知') 'SupCode'"`
	BrandCode   string `xorm:"comment('品牌') 'BrandCode'"`
	XTaxRate    string `xorm:"comment('税率') 'XTaxRate'"`
	IsWeight    string `xorm:"comment('是否称重') 'IsWeight'"`
	PluStatus   string `xorm:"comment('是否称重') 'PluStatus'"`
	Produce     string `xorm:"comment('税收分类编码') 'Produce'"`
	Grade       string `xorm:"comment('税收分类编码简称') 'Grade'"`
	XgDate      string `xorm:"comment('更新日期') 'XgDate'"`
	LrDate      string `xorm:"comment('创建日期') 'LrDate'"`
}

// Hander 数据预处理
func (i *Item) Hander() {
	i.PluCode = util.TrimSpace(i.PluCode)
	i.PluName = util.TrimSpace(gbk.Decode(i.PluName))
	i.PluAbbrName = util.TrimSpace(gbk.Decode(i.PluAbbrName))
	i.BarCode = util.TrimSpace(i.BarCode)
	i.Spec = util.TrimSpace(gbk.Decode(i.Spec))
	i.SPrice = util.TrimSpace(i.SPrice)
	i.HJPrice = util.TrimSpace(i.HJPrice)
	i.Unit = util.TrimSpace(i.Unit)
	i.DepCode = util.TrimSpace(i.DepCode)
	i.ClsCode = util.TrimSpace(i.ClsCode)
	i.SupCode = util.TrimSpace(i.SupCode)
	i.BrandCode = util.TrimSpace(i.BrandCode)
	i.XTaxRate = util.TrimSpace(i.XTaxRate)
	i.IsWeight = util.TrimSpace(i.IsWeight)
	i.PluStatus = util.TrimSpace(i.PluStatus)
	i.Produce = util.TrimSpace(i.Produce)
	i.Grade = util.TrimSpace(gbk.Decode(i.Grade))
}
