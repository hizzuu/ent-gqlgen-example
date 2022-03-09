package utils

import "github.com/hizzuu/plate-backend/conf"

func ReadConf() {
	conf.ReadConfig(conf.ReadConfigOption{Env: "test"})
}
