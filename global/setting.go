package global

import "sqlStruct/pkg/setting"

type MysqlConnectPool struct {
}

var (
	DatabaseSetting *setting.DatabaseSettingS
	ModelPath  string
	ModelReplace string
	MysqlInstance *MysqlConnectPool
)
