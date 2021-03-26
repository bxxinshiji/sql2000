package service

import (
	"strconv"
	"strings"
	"time"

	pd "github.com/bxxinshiji/sql2000/proto/department"
	"github.com/bxxinshiji/sql2000/service/repository/models"
	"github.com/go-xorm/xorm"
	"github.com/shopspring/decimal"
)

//Department 商品仓库接口
type Department interface {
	Sale(req *pd.Request) (int64, error)
}

// DepartmentRepository 用户仓库
type DepartmentRepository struct {
	Engine *xorm.Engine
	Engine1 *xorm.Engine
}

// Sale 获取日报表总和
func (repo *DepartmentRepository) Sale(req *pd.Request) (int64, error) {
	var engine  *xorm.Engine
	switch req.Database {
	case "boxing":
		engine = srv.Engine
	case "chunliang":
		engine = srv.Engine1
	default:
		return false, nil, fmt.Errorf("database empty")
	}
	Start, _ := time.Parse("2006-01-02T15:04:05+08:00", req.StartDate)
	End, _ := time.Parse("2006-01-02T15:04:05+08:00", req.EndDate)
	End = End.Add(-1) // 修正跨年 bug 结束时间减去1
	// 获取查询年份
	// h, _ := time.ParseDuration("1h")
	startYear := Start.Year()
	sql := "SELECT sum(JxSsAmt+DxSsAmt+LxSsAmt) as Money FROM tRptDepSale" + strconv.Itoa(startYear)
	whereParts := make([]string, 0)
	whereParts = append(whereParts, `RptDate >= '`+Start.Format("2006-01-02")+`'`)
	whereParts = append(whereParts, `RptDate <= '`+End.Format("2006-01-02")+`'`)
	if len(whereParts) > 0 {
		sql = sql + " WHERE " + strings.Join(whereParts, " AND ")
	}
	whereOrParts := make([]string, 0)
	for _, dep := range req.Department {
		whereOrParts = append(whereOrParts, `DepCode = '`+strconv.FormatInt(dep, 10)+`'`)
	}
	if len(whereOrParts) > 0 {
		sql = sql + " AND (" + strings.Join(whereOrParts, " OR ") + ")"
	}
	if req.Where != "" {
		sql = sql + " AND " + req.Where
	}
	// fmt.Println(sql)
	dep := &models.Department{}
	_, err := engine.SQL(sql).Get(dep)
	if err != nil {
		return 0, err
	}
	// 乘以 100 返回 int64
	total := dep.Money.Mul(decimal.NewFromFloat(100)).IntPart()

	endYear := End.Year()
	// 多年查询的时候累加 分数据库查询
	for i := startYear; i < endYear; i++ {
		sql = strings.Replace(sql, "tRptDepSale"+strconv.Itoa(i), "tRptDepSale"+strconv.Itoa(i+1), -1)
		_, err := engine.SQL(sql).Get(dep)
		if err != nil {
			return 0, err
		}
		// 两个数据库相加
		total = total + dep.Money.Mul(decimal.NewFromFloat(100)).IntPart()
	}
	return total, err
}
