package handler

import (
	"context"
	"encoding/json"
	"fmt"

	cli "github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/util/log"

	client "github.com/lecex/core/client"
	"github.com/lecex/core/env"

	service "github.com/bxxinshiji/sql2000/service/repository"
)

// Sync 同步
type Sync struct {
	CashierRepo service.Cashier
}

// Get 获取商品详细
func (srv *Sync) Cashier() {
	users, err := srv.CashierRepo.All()
	if err != nil {
		log.Fatal(err)
	}
	valid, err := srv.cashierExist(users[0]["Code"])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(123, users, valid)
}

func (srv *Sync) cashierExist(code string) (valid bool, err error) {
	req := make(map[string]interface{})
	err = json.Unmarshal([]byte(`{
		"cashier": {
			"code": "`+code+`"
		}
	}`), &req)
	if err != nil {
		return false, err
	}
	res := make(map[string]interface{})
	err = client.Call(context.TODO(), env.Getenv("MICRO_API_NAMESPACE", "go.micro.api.")+"device", "Cashiers.Exist", &req, &res, cli.WithContentType("application/json"))
	if err != nil {
		return false, err
	}
	if valid, ok := res["valid"]; ok {
		if valid.(bool) {
			return true, nil
		} else {
			return false, nil
		}
	} else {
		return false, nil
	}
}
