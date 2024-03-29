package config

import (
	"github.com/lecex/core/config"
	"github.com/lecex/core/env"
)

// Conf 配置
var Conf config.Config = config.Config{
	Name:    env.Getenv("MICRO_API_NAMESPACE", "go.micro.srv.") + "sql2000",
	Version: "v1.3.3",
}
