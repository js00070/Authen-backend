package model

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB 数据库连接
var DB *gorm.DB

func migration() {
	// 自动迁移模式
	err := DB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=UTF8MB4").AutoMigrate(
		&User{}, &File{}).Error
	if err != nil {
		panic(err)
	}
}

// Init 初始化数据库连接
func Init(connStr string) {
	log.Println(connStr)
	db, err := gorm.Open("mysql", connStr)
	db.LogMode(true)
	// Error
	if err != nil {
		panic(err)
	}
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(20)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db
	migration()
}
