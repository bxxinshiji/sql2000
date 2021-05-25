package service

import (
	"fmt"
	"strconv"

	"github.com/MXi4oyu/Utils/cnencoder/gbk"
	"github.com/bxxinshiji/sql2000/util"

	pd "github.com/bxxinshiji/sql2000/proto/item"
	"github.com/bxxinshiji/sql2000/service/repository/models"
	"github.com/go-xorm/xorm"
	"github.com/goinggo/mapstructure"
)

//Item 商品仓库接口
type Item interface {
	All() (*pd.Goods, error)
	Get(item *pd.Item, database string) (bool, *pd.Item, error)
}

// ItemRepository 商品仓库
type ItemRepository struct {
	Engine  *xorm.Engine
	Engine1 *xorm.Engine
}

//AllGet 获取所有简易商品信息
func (srv *ItemRepository) All() (goods *pd.Goods, err error) {
	goods = &pd.Goods{
		Plus: make(map[string]*pd.Info),
		Bars: make(map[string]*pd.Info),
	}
	sqlPlu := `select
			PluCode,
			BarCode,
			PluName as Name,
			SPrice as Price
		from tBmPlu
		WHERE	PluStatus='1'`
	resPlu, err := srv.Engine.QueryString(sqlPlu)
	for _, item := range resPlu {
		price, err := strconv.ParseFloat(util.TrimSpace(item["Price"]), 64)
		if err != nil {
			return nil, err
		}
		pluCode := util.TrimSpace(item["PluCode"])
		goods.Plus[pluCode] = &pd.Info{
			PluCode: pluCode,
			BarCode: util.TrimSpace(item["BarCode"]),
			Name:    util.TrimSpace(gbk.Decode(item["Name"])),
			Price:   int64(price * 100),
		}
	}
	sqlBar := `select 
            a.PluCode as PluCode,
			a.BarCode as BarCode,
			a.PluName as Name,
			b.SPrice as Price
		from tbmMulBar as a LEFT JOIN tBmPlu b ON a.PluCode=b.PluCode 
		WHERE b.PluStatus='1'`
	resBar, err := srv.Engine.QueryString(sqlBar)
	for _, item := range resBar {
		price, err := strconv.ParseFloat(util.TrimSpace(item["Price"]), 64)
		if err != nil {
			return nil, err
		}
		pluCode := util.TrimSpace(item["PluCode"])
		goods.Bars[pluCode] = &pd.Info{
			PluCode: pluCode,
			BarCode: util.TrimSpace(item["BarCode"]),
			Name:    util.TrimSpace(gbk.Decode(item["Name"])),
			Price:   int64(price * 100),
		}
	}
	return
}

// Get 获取商品信息
func (srv *ItemRepository) Get(item *pd.Item, database string) (bool, *pd.Item, error) {
	var engine *xorm.Engine
	switch database {
	case "boxing":
		engine = srv.Engine
	case "chunliang":
		engine = srv.Engine1
	default:
		return false, nil, fmt.Errorf("database empty")
	}
	itemModel := &models.Item{}
	if item.BarCode != "" {
		itemModel.BarCode = item.BarCode
	}
	if item.PluCode != "" {
		itemModel.PluCode = item.PluCode
	}
	res, bars, err := srv.itemInfo(engine, itemModel)
	if err != nil || !res {
		return false, nil, err
	}
	item.Bars = bars
	item.PluCode = itemModel.PluCode
	item.BarCode = itemModel.BarCode
	item.Name = itemModel.PluName
	item.Price = itemModel.SPrice
	item.Status = itemModel.PluStatus

	item.BuyPrice = itemModel.HJPrice
	item.Spec = itemModel.Spec
	item.Unit = itemModel.Unit
	item.Type = itemModel.IsWeight
	item.DeptCode = itemModel.DepCode
	item.BrandCode = itemModel.BrandCode
	item.CreatedAt = itemModel.LrDate
	item.UpdatedAt = itemModel.XgDate
	stock, err := srv.itemStock(engine, item.PluCode)
	if err != nil {
		return false, nil, err
	}
	item.Stock = stock

	return true, item, err
}

// itemStock 商品库存
func (srv *ItemRepository) itemStock(engine *xorm.Engine, pluCode string) (stock *pd.Stock, err error) {
	stock = &pd.Stock{}
	supplier := make([]*pd.Supplier, 0)
	sql := `
		select 
            k.SupCode as SupCode,
            k.PluCode as PluCode,
            (k.KcJxNumber + k.KcDxNumber) as Number,
			s.SupName as Name
    	from tYwPluKcSup as k, tBmSup as s
    	WHERE k.PluCode = '` + pluCode + `' AND s.SupCode = k.SupCode 
    	ORDER BY k.SupCode, k.DepCode
	`
	results, err := engine.QueryString(sql)
	if err != nil {
		return stock, err
	}
	supStock := &models.SupStock{}
	for _, res := range results {
		err := mapstructure.Decode(res, supStock)
		supStock.Hander()
		if err != nil {
			return stock, err
		}
		supplier = append(supplier, &pd.Supplier{
			Code:   supStock.SupCode,
			Number: supStock.Number,
			Name:   supStock.Name,
		})
	}
	stock.Supplier = supplier
	return stock, err
}

// ItemInfo 商品信息
func (srv *ItemRepository) itemInfo(engine *xorm.Engine, item *models.Item) (res bool, bars []*pd.Bar, err error) {
	res, err = engine.Table("tBmPlu").Get(item)
	if err != nil {
		return false, bars, err
	}
	if !res && item.PluCode != "" { // plucode查询未果返回
		return res, bars, err
	}
	if !res && item.BarCode != "" {
		// 多条形码时获取指定商品ID
		barCode := &models.BarItem{
			BarCode: item.BarCode,
		}
		res, err = engine.Table("tbmMulBar").Get(barCode)
		if err != nil {
			return false, bars, err
		}
		// 重新获取商品信息
		if res {
			// 缓存原始条形码
			item.BarCode = ""
			item.PluCode = barCode.PluCode
			_, err := engine.Table("tBmPlu").Get(item)
			if err != nil {
				return false, bars, err
			}
		}
		if !res && item.BarCode != "" { // plucode查询未果返回
			return res, bars, err
		}
	}
	if item.PluCode != "" { // 获取多条码数据
		bar := new(models.BarItem)
		rows, err := engine.Table("tbmMulBar").Where("PluCode = ?", item.PluCode).Rows(bar)
		if err != nil {
			return false, bars, err
		}
		defer rows.Close()
		for rows.Next() {
			err = rows.Scan(bar)
			bar.Hander() // 处理返回数据
			bars = append(bars, &pd.Bar{
				BarCode: bar.BarCode,
				PluCode: bar.PluCode,
				Name:    bar.PluName,
				Spec:    bar.Spec,
			})
		}
	}
	// 处理梳理
	item.Hander()
	return true, bars, err
}
