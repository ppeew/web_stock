package model

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 写数据库内模型
var DB *gorm.DB

type User struct {
	gorm.Model
	Name   string
	Passwd string
	Email  string
}

func init() {
	//打开数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/web_stock?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}
	DB = d
	err2 := DB.AutoMigrate(&User{})
	if err2 == nil {
		fmt.Println("添加成功")
	}
}
