package database

import (
	mercTypeData "merchandise-circulation-api/src/app/merchandise_types/data"
	mercData "merchandise-circulation-api/src/app/merchandises/data"
	userData "merchandise-circulation-api/src/app/users/repositories"
	"merchandise-circulation-api/src/configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var DB *gorm.DB
type DBConf struct{ *gorm.DB }

func (DB *DBConf) InitDB() *DBConf {
	config, _ := configs.LoadServerConfig(".")
	dsn := config.ConnectionString

	conn, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return &DBConf{conn}
}

func (DB *DBConf) Migrate() {
	DB.AutoMigrate(&mercData.Merchandise{}, &mercTypeData.MerchandiseType{},
		&userData.User{})
}

func (DB *DBConf) Demigrate() {
	DB.Migrator().DropTable(&mercData.Merchandise{},
		&mercTypeData.MerchandiseType{}, &userData.User{})
}
