package databases

import (
	"go-auth/configs/databases/mysqlConf"
	"go-auth/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connetct() {
	conn, err := gorm.Open(mysql.Open(mysqlConf.New()), &gorm.Config{})

	if err != nil {
		panic("could not connetct to the mysql database")
	}

	DB = conn

	conn.AutoMigrate(&models.User{})
}
