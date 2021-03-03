package service

import (
	pd "github.com/bxxinshiji/sql2000/proto/item"
	"github.com/bxxinshiji/sql2000/service/repository/models"
	"github.com/go-xorm/xorm"
)

//Item 商品仓库接口
type Item interface {
	Get(item *pd.Item) (bool, *pd.Item, error)
}

// ItemRepository 商品仓库
type ItemRepository struct {
	Engine *xorm.Engine
}

// Get 获取商品信息
func (srv *ItemRepository) Get(item *pd.Item) (bool, *pd.Item, error) {
	itemModel := &models.Item{}
	if item.BarCode != "" {
		itemModel.BarCode = item.BarCode
	}
	if item.PluCode != "" {
		itemModel.PluCode = item.PluCode
	}
	res, bars, err := srv.itemInfo(srv.Engine, itemModel)
	if err != nil || !res {
		return false, nil, err
	}
	item.Bars = bars
	item.PluCode = itemModel.PluCode
	item.BarCode = itemModel.BarCode
	item.Name = itemModel.PluName
	item.Price = itemModel.SPrice
	item.BuyPrice = itemModel.HJPrice
	item.Spec = itemModel.Spec
	item.Unit = itemModel.Unit
	item.Type = itemModel.IsWeight
	item.Status = itemModel.PluStatus
	item.DeptCode = itemModel.DepCode
	item.BrandCode = itemModel.BrandCode
	item.CreatedAt = itemModel.LrDate
	item.UpdatedAt = itemModel.XgDate
	return true, item, err
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
