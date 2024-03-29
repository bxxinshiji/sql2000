package handler

import (
	"context"

	pb "github.com/bxxinshiji/sql2000/proto/item"
	service "github.com/bxxinshiji/sql2000/service/repository"
)

// Item 商品结构
type Item struct {
	Repo service.Item
}

// Get 获取商品详细
func (srv *Item) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	valid, item, err := srv.Repo.Get(req.Item, req.Database)
	if err != nil {
		return err
	}
	res.Valid = valid
	res.Item = item
	return err
}

// All 获取所有简易商品详细
func (srv *Item) All(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	goods, err := srv.Repo.All()
	if err != nil {
		return err
	}
	res.Goods = goods
	return err
}
