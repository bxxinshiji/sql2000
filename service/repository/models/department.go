package models

import (
	"time"

	"github.com/bxxinshiji/sql2000/util"
	"github.com/shopspring/decimal"
)

// Department 商品
type Department struct {
	Date       time.Time       `xorm:"comment('结算时间') 'RptDate'"`
	Department string          `xorm:"comment('部门编码') 'DepCode'"`
	Money      decimal.Decimal `xorm:"comment('实收金额') 'Money'"`
	// 三个字段 JxSsAmt、DxSsAmt、LxSsAmt
}

// Hander 数据预处理
func (p *Department) Hander() {
	p.Department = util.TrimSpace(p.Department)
}
