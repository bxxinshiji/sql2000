package handler

import (
	"context"
	"fmt"

	pb "github.com/lecex/sql2000/proto/item"
)

// Item 商品结构
type Item struct {
}

// Get 获取商品详细
func (srv *Item) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	fmt.Println(req, res)
	return err
}
