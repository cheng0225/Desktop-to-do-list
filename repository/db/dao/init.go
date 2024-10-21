package dao

import (
	"context"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"schedule/config"
)

var (
	_db *gorm.DB
)

func InitSqlite() {
	// 初始化数据库
	// fmt.Println(config.DBFile)
	db, err := gorm.Open(sqlite.Open(config.DBFile), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	_db = db
	// 自动迁移创建表
	Migration()
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
