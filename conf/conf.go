package conf

import (
	"path/filepath"
	"strings"

	"github.com/hizzuu/plate-backend/utils/path"
	"github.com/spf13/viper"
)

type ReadConfigOption struct {
	Env string
}

type conf struct {
	App app
	DB  db
	Api api
}

type app struct {
	Debug bool
	Env   string
}

type db struct {
	Dbms                 string
	User                 string
	Pass                 string
	Net                  string
	Host                 string
	Port                 string
	Name                 string
	ParseTime            bool
	AllowNativePasswords bool
}

type api struct {
	Port            string
	StorageName     string `mapstructure:"storage_name"`
	FirebaseKeyJson string `mapstructure:"firebase_key_json"`
}

var C *conf

func ReadConfig(option ReadConfigOption) {
	switch option.Env {
	case "test":
		viper.SetConfigName("conf.test")
	default:
		viper.SetConfigName("conf")
	}

	viper.AddConfigPath(filepath.Join(path.RootDir(), "/conf"))
	viper.SetConfigType("yml")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&C); err != nil {
		panic(err)
	}
}
