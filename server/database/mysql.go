package database

import (
    "fmt"
    "gin/model"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

// Migrate
func Migrate()  {
    // 自动迁移
    InitDb()
    DB.AutoMigrate(&model.Url{})
    fmt.Println("")
}
func InitDb(){
    dsn := "root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
    DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
