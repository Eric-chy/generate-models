package dbtools

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"sqlStruct/global"
	"sync"
)

/*
* MysqlConnectiPool
* 数据库连接操作库
* 基于gorm封装开发
 */
type MysqlConnectPool struct {
}

var mysqlInstance *MysqlConnectPool
var mysqlOnce sync.Once


var db *gorm.DB
var err_db error


func GetMysqlInstance() *MysqlConnectPool {
	mysqlOnce.Do(func() {
		mysqlInstance = &MysqlConnectPool{}
	})
	return mysqlInstance
}

/*
* @fuc 初始化数据库连接(可在mail()适当位置调用)
 */
func (m *MysqlConnectPool) InitMysqlPool() (issucc bool) {
	db, err_db = gorm.Open("mysql", global.DatabaseSetting.UserName+":"+global.DatabaseSetting.Password+"@tcp("+global.DatabaseSetting.Host +")/"+global.DatabaseSetting.DBName+"?charset=utf8&parseTime=True&loc=Local")
	db.SingularTable(true)
	if err_db != nil {
		log.Fatal(err_db)
		return false
	}
	//关闭数据库，db会被多个goroutine共享，可以不调用
	// defer db.Close()
	return true
}

/*
* @fuc  对外获取数据库连接对象db
 */
func (m *MysqlConnectPool) GetMysqlPool() *gorm.DB {
	//db.LogMode(true)
	return db
}

func GetMysqlDb() (db *gorm.DB) {
	return GetMysqlInstance().GetMysqlPool()
}




