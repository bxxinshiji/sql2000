package main

import (
	"testing"

	db "github.com/bxxinshiji/sql2000/providers/database"
	service "github.com/bxxinshiji/sql2000/service/repository"

	"github.com/bxxinshiji/sql2000/handler"
)

// func TestItemsGet(t *testing.T) {
// 	req := &itemPB.Request{
// 		Item: &itemPB.Item{
// 			BarCode: `6923450662007`,
// 		},
// 		Database: "chunliang",
// 	}
// 	res := &itemPB.Response{}
// 	h := handler.Item{&service.ItemRepository{db.Engine,db.Engine1}}
// 	err := h.Get(context.TODO(), req, res)
// 	fmt.Println("ItemGet", res, err)
// 	t.Log(req, res, err)
// }

func TestItemsAll(t *testing.T) {
	// req := &itemPB.Request{}
	// res := &itemPB.Response{}
	// h := handler.Item{&service.ItemRepository{db.Engine, db.Engine1}}
	// err := h.All(context.TODO(), req, res)
	// fmt.Println("ItemGet", res, err)
	// t.Log(req, res, err)
}

func TestSync(t *testing.T) {
	sync := handler.Sync{CashierRepo: &service.CashierRepository{db.Engine, db.Engine1}}
	sync.Cashier()
}
