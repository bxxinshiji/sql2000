package service

import (
	"github.com/go-xorm/xorm"

	"github.com/MXi4oyu/Utils/cnencoder/gbk"
	"github.com/bxxinshiji/sql2000/util"
)

//Cashier 商品仓库接口
type Cashier interface {
	All() ([]map[string]string, error)
}

// CashierRepository 商品仓库
type CashierRepository struct {
	Engine  *xorm.Engine
	Engine1 *xorm.Engine
}

//All 获取所有简易商品信息
func (srv *CashierRepository) All() ([]map[string]string, error) {
	sql := `select
			UserCode as Code,
			UserName as Name,
			Passwd as Password
		from tXsUser`
	res, err := srv.Engine.QueryString(sql)
	for _, item := range res {
		item["Code"] = util.TrimSpace(item["Code"])
		item["Name"] = util.TrimSpace(gbk.Decode(item["Name"]))
		item["Password"] = util.TrimSpace(item["Password"])
	}
	return res, err
}
