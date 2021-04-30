package global

import "github.com/jinzhu/gorm"

type MysqlConnectPool struct {
}

var (
	ModelPath  string
	ModelReplace string
	DBEngine    *gorm.DB
)
