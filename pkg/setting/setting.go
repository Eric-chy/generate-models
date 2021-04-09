package setting

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting(cfg, file string) (*Setting, error) {
	_, err := os.Stat(cfg + file + ".yaml")
	if os.IsNotExist(err) {
		fmt.Println(cfg + file + "文件不存在")
		return nil, err
	}
	vp := viper.New()
	vp.SetConfigName(file)
	vp.AddConfigPath(cfg)
	//vp.SetConfigType("yaml")
	err = vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Setting{vp}, nil
}
