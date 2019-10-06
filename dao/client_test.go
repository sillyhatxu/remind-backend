package dao

import (
	"fmt"
	"github.com/sillyhatxu/environment-config"
)

func init() {
	path := "/Users/shikuanxu/go/src/github.com/sillyhatxu/remind-backend"
	err := envconfig.LoadEnv(fmt.Sprintf("%s/%s", path, ".env"))
	if err != nil {
		panic(err)
	}
	err = envconfig.ParseDefaultConfig(envconfig.ConfigFile(fmt.Sprintf("%s/%s", path, "config-local.conf")))
	if err != nil {
		panic(err)
	}
	err = InitialMysqlClient()
	if err != nil {
		panic(err)
	}
}
