package database

import (
	"log"

	"github.com/go-xorm/xorm"
	"github.com/lecex/core/env"
	txorm "github.com/lecex/core/xorm"
)

// Engine 管理包
var (
	// Engine1 连接
	Engine  *xorm.Engine
	Engine1 *xorm.Engine
)

func init() {
	var err error
	conf := &txorm.Config{
		Driver: env.Getenv("DB_DRIVER", "odbc"),
		// Host 主机地址
		Host: env.Getenv("DB_HOST", "192.168.5.10"),
		// Port 主机端口
		Port: env.Getenv("DB_PORT", "1433"),
		// User 用户名
		User: env.Getenv("DB_USER", "sa"),
		// Password 密码
		Password: env.Getenv("DB_PASSWORD", "123456"),
		// DbName 数据库名称
		DbName: env.Getenv("DB_NAME", "stmis1"),
		// Charset 数据库编码
		Charset: env.Getenv("DB_CHARSET", "GBK"),
	}
	Engine, err = txorm.Connection(conf)
	if err != nil {
		log.Fatalf("connect error: %v\n", err)
	}
	conf1 := &txorm.Config{
		Driver: env.Getenv("DB_DRIVER_1", "odbc"),
		// Host 主机地址
		Host: env.Getenv("DB_HOST_1", "192.168.20.10"),
		// Port 主机端口
		Port: env.Getenv("DB_PORT_1", "1433"),
		// User 用户名
		User: env.Getenv("DB_USER_1", "sa"),
		// Password 密码
		Password: env.Getenv("DB_PASSWORD_1", "123456"),
		// DbName 数据库名称
		DbName: env.Getenv("DB_NAME_1", "stmis1"),
		// Charset 数据库编码
		Charset: env.Getenv("DB_CHARSET_1", "GBK"),
	}
	Engine1, err = txorm.Connection(conf1)
	if err != nil {
		log.Fatalf("connect error: %v\n", err)
	}
}
