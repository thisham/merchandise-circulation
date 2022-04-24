package database

import (
	mercTypeData "merchandise-circulation-api/src/app/merchandise_types/data"
	mercData "merchandise-circulation-api/src/app/merchandises/data"
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

func Migrate() {
	InitDB()
	DB.AutoMigrate(mercData.Merchandise{}, mercTypeData.MerchandiseType{})
}

func Demigrate() {
	InitDB()
	DB.Migrator().DropTable(mercData.Merchandise{}, mercTypeData.MerchandiseType{})
}
