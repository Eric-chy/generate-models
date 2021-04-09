package main

import (
	"flag"
	"fmt"
	"os"
	"sqlStruct/dbtools"
	"sqlStruct/generate"
	. "sqlStruct/global"
	"sqlStruct/pkg/mysql"
	"sqlStruct/pkg/setting"
	"strings"
)

var (
	cfg    string
	file   string
	path   string
	rp     string
	db     string
	tables string
)

func main() {
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println(err)
	//	}
	//}()
	boot()
	//初始化数据库
	dbtools.Init()
	mysql.Init()
	generate.Genertate(DatabaseSetting.Tables) //生成指定表信息，可变参数可传入多个表名
}

func boot() {
	flag.StringVar(&cfg, "cfg", "configs/", "指定要使用的配置文件路径")
	flag.StringVar(&file, "file", "dev.config", "指定要使用的配置文件名")
	flag.StringVar(&path, "path", "./internal/model/", "指定要生成的model路径")
	flag.StringVar(&rp, "rp", "n", "是否替换")
	flag.StringVar(&db, "db", "", "数据库名")
	flag.StringVar(&tables, "tables", "", "表名")
	flag.Parse()
	s, err := setting.NewSetting(cfg, file)
	if err != nil {
		fmt.Println("读取配置文件失败", err)
		os.Exit(500)
	}
	err = s.ReadSection("Database", &DatabaseSetting)
	if db != "" {
		DatabaseSetting.DBName = db
	}
	if tables != "" {
		DatabaseSetting.Tables = strings.Split(tables, ",")
	}
	if err != nil {
		fmt.Println("读取mysql配置失败", err)
		os.Exit(500)
	}
	//生成model的文件夹
	ModelPath = path
	//判断生成model的文件夹是否存在，不存在则递归创建
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		err := os.MkdirAll(path, 0777)
		if err != nil {
			fmt.Printf("model目录不存在")
		}
	}
	//是否覆盖生成
	ModelReplace = rp
}
