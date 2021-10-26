package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday	 *time.Time
}

func main()  {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// 自动迁移
	db.AutoMigrate(&User{})

	//user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	//database.Create(&user)
	//var user User
	//database.Where("name = ?", "jinzhu").First(&user)
	//print(user.Age)
}