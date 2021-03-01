package service

import (
	"github.com/go-xorm/xorm"
	pd "github.com/bxxinshiji/sql2000/proto/item"
	"github.com/bxxinshiji/sql2000/service/repository/models"
)

//Item 商品仓库接口
type Item interface {
	Get(item *pd.Item) (*pd.Item, error)
}

// ItemRepository 商品仓库
type ItemRepository struct {
	Engine *xorm.Engine
}

// Get 获取商品信息
func (srv *ItemRepository) Get(item *pd.Item) (*pd.Item, error) {
	itemModel := &models.Item{}
	if item.BarCode != "" {
		itemModel.BarCode = item.BarCode
	}
	if item.PluCode != "" {
		itemModel.PluCode = item.PluCode
	}
	bars, err := srv.itemInfo(srv.Engine, itemModel)
	if err != nil {
		return nil, err
	}
	item.Bars = bars
	item.PluCode = itemModel.PluCode
	item.BarCode = itemModel.BarCode
	item.Name = itemModel.PluName
	item.Price = itemModel.SPrice
	item.Spec = itemModel.Spec
	item.Unit = itemModel.Unit
	item.Type = itemModel.IsWeight
	item.Status = itemModel.PluStatus
	item.DeptCode = itemModel.DepCode
	item.BrandCode = itemModel.BrandCode
	item.CreatedAt = itemModel.LrDate
	item.UpdatedAt = itemModel.XgDate
	return item, err
}

// ItemInfo 商品信息
func (srv *ItemRepository) itemInfo(engine *xorm.Engine, item *models.Item) (bars []*pd.Bar, err error) {
	res, err := engine.Table("tBmPlu").Get(item)
	if err != nil {
		return bars, err
	}
	if !res && item.BarCode != "" {
		// 多条形码时获取指定商品ID
		barCode := &models.BarItem{
			BarCode: item.BarCode,
		}
		res, err := engine.Table("tbmMulBar").Get(barCode)
		if err != nil {
			return bars, err
		}
		// 重新获取商品信息
		if res {
			// 缓存原始条形码
			item.BarCode = ""
			item.PluCode = barCode.PluCode
			_, err := engine.Table("tBmPlu").Get(item)
			if err != nil {
				return bars, err
			}
		}
	}
	if item.PluCode != "" { // 获取多条码数据
		bar := new(models.BarItem)
		rows, err := engine.Table("tbmMulBar").Where("PluCode = ?", item.PluCode).Rows(bar)
		if err != nil {
			return bars, err
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
	return bars, err
}
