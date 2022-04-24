package database

import (
	"merchandise-circulation-api/src/configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	config, _ := configs.LoadServerConfig(".")
	dsn := config.ConnectionString

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
