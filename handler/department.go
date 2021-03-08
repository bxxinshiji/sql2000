package handler

import (
	"context"

	pb "github.com/bxxinshiji/sql2000/proto/department"
	service "github.com/bxxinshiji/sql2000/service/repository"
)

// Department 部门报表
type Department struct {
	Dep service.Department
}

// Sale 部门收款金额总和
func (srv *Department) Sale(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	total, err := srv.Dep.Sale(req)
	if err != nil {
		return err
	}
	res.Total = total
	return err
}
