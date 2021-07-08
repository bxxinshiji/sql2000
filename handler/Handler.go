package handler

import (
	"time"

	server "github.com/micro/go-micro/v2/server"

	departmentPB "github.com/bxxinshiji/sql2000/proto/department"
	itemPB "github.com/bxxinshiji/sql2000/proto/item"

	db "github.com/bxxinshiji/sql2000/providers/database"
	service "github.com/bxxinshiji/sql2000/service/repository"
)

// Register 注册
func Register(Server server.Server) {
	itemPB.RegisterItemsHandler(Server, &Item{&service.ItemRepository{db.Engine, db.Engine1}}) // 用户服务实现                                       // 权限管理服务实现
	departmentPB.RegisterDepartmentHandler(Server, &Department{&service.DepartmentRepository{db.Engine, db.Engine1}})

	go sync() // 同步sql2000数据
}

// sync 同步
func sync() {
	time.Sleep(30 * time.Second)
	sync := Sync{CashierRepo: &service.CashierRepository{db.Engine, db.Engine1}}
	sync.Cashier()
}
