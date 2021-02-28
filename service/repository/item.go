package service

import (
	"github.com/go-xorm/xorm"
	pd "github.com/lecex/sql2000/proto/item"
	"github.com/lecex/sql2000/service/repository/models"
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
	itemModel := &models.Item{
		BarCode: item.BarCode,
	}
	err := srv.itemInfo(srv.Engine, itemModel)
	if err != nil {
		return nil, err
	}
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
func (srv *ItemRepository) itemInfo(engine *xorm.Engine, item *models.Item) (err error) {
	res, err := engine.Table("tBmPlu").Get(item)
	if err != nil {
		return err
	}
	if !res {
		// 多条形码时获取指定商品ID
		barCode := &models.BarItem{
			BarCode: item.BarCode,
		}
		res, err := engine.Table("tbmMulBar").Get(barCode)
		if err != nil {
			return err
		}
		// 重新获取商品信息
		if res {
			// 缓存原始条形码
			item.BarCode = ""
			item.PluCode = barCode.PluCode
			_, err := engine.Table("tBmPlu").Get(item)
			if err != nil {
				return err
			}
			item.BarCode = barCode.BarCode
			item.DepCode = barCode.DepCode
			item.PluName = barCode.PluName
			item.Spec = barCode.Spec
		}
	}
	// 处理梳理
	item.Hander()
	return err
}
