package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"sqlStruct/global"
	"os"
	"sync"
)

/*
* MysqlConnectiPool
* 数据库连接操作库
* 基于gorm封装开发
 */
type Pool struct {
}

var (
	Singleton *Pool
	once sync.Once
	db *gorm.DB
	err error
)

func Init(){
	//初始化Mysql连接池
	conn := GetSingleton().initPool()
	if !conn {
		log.Println("init database pool failure...")
		os.Exit(1)
	}
}

func (p *Pool) initPool() bool {
	db, err = gorm.Open("mysql", global.DatabaseSetting.UserName+":"+global.DatabaseSetting.Password+"@tcp("+global.DatabaseSetting.Host +")/"+global.DatabaseSetting.DBName+"?charset=utf8&parseTime=True&loc=Local")
	db.SingularTable(true)
	if err != nil {
		log.Fatal(err)
		return false
	}
	// defer db.Close()
	return true
}

func GetSingleton() *Pool {
	once.Do(func() {
		Singleton = &Pool{}
	})
	return Singleton
}

func (p *Pool) Get() *gorm.DB {
	return db
}

func GetDb() (db *gorm.DB) {
	return GetSingleton().Get()
}



