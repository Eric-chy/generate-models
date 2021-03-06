package main

import (
	"flag"
	"fmt"
	"generate-models/configs"
	"generate-models/generate"
	. "generate-models/global"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
	"strings"
)

var (
	cfgPath    string
	file   string
	modelPath   string
	replace  string
	db     string
	tables string
)

func main() {
	flagParse()
	setupDb()
	//初始化数据库
	generate.Genertate(configs.Conf.Database.Tables) //生成指定表信息，可变参数可传入多个表名
}

func flagParse() {
	flag.StringVar(&cfgPath, "c", "configs/", "指定要使用的配置文件路径")
	flag.StringVar(&file, "f", "dev", "指定要使用的配置文件名")
	flag.StringVar(&modelPath, "m", "./models/", "指定要生成的model路径")
	flag.StringVar(&replace, "r", "n", "是否替换, 默认n，可选y|n")
	flag.StringVar(&db, "d", "", "数据库名,不填则按配置文件来")
	flag.StringVar(&tables, "t", "", "表名,多个表用英文半角,隔开")
	if strings.ToLower(replace) != "n" && strings.ToLower(replace) != "y" {
		replace = "n"
	}
	flag.Parse()
	configs.Init(cfgPath, file)
	if db != "" {
		configs.Conf.Database.DBName = db
	}
	if tables != "" {
		configs.Conf.Database.Tables = strings.Split(tables, ",")
	}
	//生成model的文件夹
	ModelPath = modelPath
	//判断生成model的文件夹是否存在，不存在则递归创建
	_, err := os.Stat(modelPath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(modelPath, 0777)
		if err != nil {
			fmt.Printf("model目录不存在")
		}
	}
	//是否覆盖生成
	ModelReplace = replace
}

func setupDb() {
	var err error
	DBEngine, err = NewDBEngine()
	if err != nil {
		log.Println("初始化数据库失败", err)
		os.Exit(1)
	}
}

func NewDBEngine() (*gorm.DB, error) {
	var dsn string
	cfg := configs.Conf.Database
	switch cfg.DBType {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@%s(%s:%d)/%s",
			cfg.UserName,
			cfg.Password,
			cfg.Protocol,
			cfg.Host,
			cfg.Port,
			cfg.DBName,
		)
	default:
		log.Fatalf("invalid db driver %v\n", cfg.DBType)
	}

	db, err := gorm.Open(cfg.DBType, dsn)
	if err != nil {
		log.Fatalf("Open "+cfg.DBType+" failed. %v\n", err)
		return nil, err
	}
	db.LogMode(true)
	db.DB().SetConnMaxLifetime(cfg.MaxLifetime) //最大连接周期，超过时间的连接就close
	db.DB().SetMaxOpenConns(cfg.MaxOpenConns)   //设置最大连接数
	db.DB().SetMaxIdleConns(cfg.MaxIdleConns)   //设置闲置连接数
	//设置全局表名禁用复数
	db.SingularTable(true)

	return db, nil
}
