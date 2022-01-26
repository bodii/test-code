package database

import (
	"go-auth/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connetct() {
	conn, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/yt_go_auth?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		panic("could not connetct to the mysql database")
	}

	DB = conn

	conn.AutoMigrate(&models.User{})
}
