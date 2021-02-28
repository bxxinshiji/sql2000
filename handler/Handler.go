package handler

import (
	server "github.com/micro/go-micro/v2/server"

	itemPB "github.com/lecex/sql2000/proto/item"
)

// Register 注册
func Register(Server server.Server) {
	itemPB.RegisterItemsHandler(Server, &Item{}) // 用户服务实现                                       // 权限管理服务实现
}
