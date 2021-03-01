package main

import (
	"context"
	"fmt"
	"testing"

	itemPB "github.com/bxxinshiji/sql2000/proto/item"
	db "github.com/bxxinshiji/sql2000/providers/database"
	service "github.com/bxxinshiji/sql2000/service/repository"

	"github.com/bxxinshiji/sql2000/handler"
)

func TestItemsGet(t *testing.T) {
	req := &itemPB.Request{
		Item: &itemPB.Item{
			BarCode: `6925911516447`,
		},
	}
	res := &itemPB.Response{}
	h := handler.Item{&service.ItemRepository{db.Engine}}
	err := h.Get(context.TODO(), req, res)
	fmt.Println("ItemGet", res, err)
	t.Log(req, res, err)
}
