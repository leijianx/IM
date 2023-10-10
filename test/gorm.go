package mq

import (
	"fmt"
	"ginchat/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func UserTest() {
	dsn := "root:secret@tcp(154.8.197.123:23306)/ginchat?charset=utf8mb4&parseTime=true&loc=Local"

	// 2. 连接服务器(池)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接失败")
		log.Fatalln(err)
	}

	// 基于模型完成表结构的迁移
	if err = db.AutoMigrate(&models.UserBasic{}); err != nil {
		log.Fatal(err)
	}
	if err = db.AutoMigrate(&models.Message{}); err != nil {
		log.Fatal(err)
	}
	if err = db.AutoMigrate(&models.Contact{}); err != nil {
		log.Fatal(err)
	}
	if err = db.AutoMigrate(&models.GroupBasic{}); err != nil {
		log.Fatal(err)
	}
	if err = db.AutoMigrate(&models.Community{}); err != nil {
		log.Fatal(err)
	}
}
