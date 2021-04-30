package configs

import (
	"github.com/spf13/viper"
	"time"
)

type Yaml struct {
	Database struct {
		DBType       string
		Protocol     string
		UserName     string
		Password     string
		Host         string
		Port         int
		DBName       string
		TablePrefix  string
		Charset      string
		ParseTime    bool
		MaxIdleConns int
		MaxOpenConns int
		Tables       []string
		MaxLifetime  time.Duration
	}
}

var Conf Yaml
var vp *viper.Viper

func Init(cfg, file string) {
	vp = viper.New()
	vp.SetConfigName(file)
	vp.AddConfigPath(cfg)
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		panic(err)
	}
	//直接整个解析
	err = vp.Unmarshal(&Conf)
	if err != nil {
		panic(err)
	}
}
