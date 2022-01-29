package main

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  sql.NullString
	Price uint
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/db_test?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}

	//定义一个表结构， 将表结构直接生成对应的表 - migrations
	// 迁移 schema
	_ = db.AutoMigrate(&Product{}) //此处应该有sql语句
	// 新增
	db.Create(&Product{Code: sql.NullString{"D42", true}, Price: 100})

}
