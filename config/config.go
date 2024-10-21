package config

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	DBFile string
)

func init() {
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
		DBFile = "notes.db"
		return
	}
	LoadSqliteData(file)
}

func LoadSqliteData(file *ini.File) {
	DBFile = file.Section("sqlite").Key("DBFile").String()
}
