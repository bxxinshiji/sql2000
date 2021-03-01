package handler

import (
	server "github.com/micro/go-micro/v2/server"

	itemPB "github.com/bxxinshiji/sql2000/proto/item"
	db "github.com/bxxinshiji/sql2000/providers/database"
	service "github.com/bxxinshiji/sql2000/service/repository"
)

// Register 注册
func Register(Server server.Server) {
	itemPB.RegisterItemsHandler(Server, &Item{&service.ItemRepository{db.Engine}}) // 用户服务实现                                       // 权限管理服务实现
}
