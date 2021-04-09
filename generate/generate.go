package generate

import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/protoc-gen-go/generator"
	"io"
	"sqlStruct/global"
	"sqlStruct/pkg/mysql"
	"os"
	"strings"
)

func Genertate(tableNames []string)  {
	//if len(tableNames) == 0 {
	//	var tables []Table
	//	dbtools.GetMysqlDb().Raw("SELECT table_name AS Name FROM information_schema.tables WHERE table_schema='"+ global.DatabaseSetting.DBName+"';").Scan(&tables)
	//	if len(tables) > 0 {
	//		for _, v := range tables {
	//			tableNames = append(tableNames, v.Name)
	//		}
	//	} else {
	//		panic(errors.New("当前库" + global.DatabaseSetting.DBName + "没有数据表"))
	//	}
	//}
	tableNamesStr := ""
	for _, name := range tableNames {
		if tableNamesStr != "" {
			tableNamesStr += ","
		}
		tableNamesStr += "'"+name+"'"
	}
	tables := getTables(tableNamesStr) //生成所有表信息
	//tables := getTables("admin_info","video_info") //生成指定表信息，可变参数可传入过个表名
	for _,table := range tables {
		fields := getFields(table.Name)
		generateModel(table,fields)
	}
}
//获取表信息
func getTables(tableNames string) []Table {
	db := mysql.GetDb()
	var tables []Table
	if tableNames == "" {
		db.Raw("SELECT TABLE_NAME as Name,TABLE_COMMENT as Comment FROM information_schema.TABLES WHERE table_schema='"+ global.DatabaseSetting.DBName+"';").Find(&tables)
	}else{
		db.Raw("SELECT TABLE_NAME as Name,TABLE_COMMENT as Comment FROM information_schema.TABLES WHERE TABLE_NAME IN ("+tableNames+") AND table_schema='"+ global.DatabaseSetting.DBName+"';").Find(&tables)
	}
	return tables
}
//获取所有字段信息
func getFields(tableName string) []Field {
	db := mysql.GetDb()
	var fields []Field
	db.Raw("show FULL COLUMNS from "+tableName+";").Find(&fields)
	return fields
}


//生成Model
func generateModel(table Table, fields []Field)  {
	content := "package model\n\n"
	//表注释
	if len(table.Comment) > 0 {
		content += "// "+table.Comment+"\n"
	}
	content += "type "+generator.CamelCase(table.Name)+" struct {\n"
	//生成字段
	for _, field := range fields {
		fieldName := generator.CamelCase(field.Field)
		fieldJson := getFieldJson(field)
		fieldType := getFiledType(field)
		fieldComment := getFieldComment(field)
		content += "	"+fieldName+" "+fieldType+" `"+fieldJson+"` "+fieldComment+"\n"
	}
	content += "}\r\n"
	first := strings.ToLower(string(generator.CamelCase(table.Name)[0]))
	content += `func (` + first + " " + generator.CamelCase(table.Name) + `) TableName() string {
	return "kbb_hotel_orders"
}`
	filename := global.ModelPath+generator.CamelCase(table.Name)+".go"
	var f *os.File
	var err error
	if checkFileIsExist(filename){
		if global.ModelReplace != "y" && global.ModelReplace != "yes" {
			fmt.Println(generator.CamelCase(table.Name)+" 已存在，需删除才能重新生成...")
			return
		}
		f, err = os.OpenFile(filename, os.O_CREATE, 0777) //打开文件
		if err != nil {
			panic(err)
		}
	}else{
		f, err = os.Create(filename)
		if err != nil {
			panic(errors.New("创建文件失败"))
		}
	}
	defer f.Close()
	_, err = io.WriteString(f, content)
	if err != nil {
		panic(err)
	}else{
		fmt.Println(generator.CamelCase(table.Name)+" 已生成...")
	}
}
//获取字段类型
func getFiledType(field Field)  string{
	typeArr := strings.Split(field.Type,"(")

	switch typeArr[0] {
	case "int":
		return "int"
	case "integer":
		return "int"
	case "mediumint":
		return "int"
	case "bit":
		return "int"
	case "year":
		return "int"
	case "smallint":
		return "int"
	case "tinyint":
		return "int"
	case "bigint":
		return "int64"
	case "decimal":
		return "float32"
	case "double":
		return "float32"
	case "float":
		return "float32"
	case "real":
		return "float32"
	case "numeric":
		return "float32"
	case "timestamp":
		return "time.Time"
	case "datetime":
		return "time.Time"
	case "time":
		return "time.Time"
	default:
		return "string"
	}
}
//获取字段json描述
func getFieldJson(field Field) string {
	return `json:"`+field.Field+`"`
}
//获取字段说明
func getFieldComment(field Field) string{
	if len(field.Comment) > 0 {
		return "// "+field.Comment
	}
	return ""
}
//检查文件是否存在
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}