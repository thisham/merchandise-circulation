package factories

import (
	"merchandise-circulation-api/src/database"

	// merchandises
	merchData "merchandise-circulation-api/src/app/merchandises/data"
	merchHandler "merchandise-circulation-api/src/app/merchandises/handlers"
	merchServices "merchandise-circulation-api/src/app/merchandises/services"

	// merchandise types
	merchTypeData "merchandise-circulation-api/src/app/merchandise_types/data"
	merchTypeHandler "merchandise-circulation-api/src/app/merchandise_types/handlers"
	merchTypeServices "merchandise-circulation-api/src/app/merchandise_types/services"

	// users
	userHandler "merchandise-circulation-api/src/app/users/handlers"
	userData "merchandise-circulation-api/src/app/users/repositories"
	userServices "merchandise-circulation-api/src/app/users/services"

	"gorm.io/gorm"
)

type handlers struct {
	MerchandiseHandler     merchHandler.MerchandiseHandler
	MerchandiseTypeHandler merchTypeHandler.MerchandiseTypeHandler
	UserHandler            userHandler.UserHandlers
}

func merchandiseFactory(conn *gorm.DB) merchHandler.MerchandiseHandler {
	data := merchData.NewMySqlRecord(conn)
	services := merchServices.NewService(data)
	return *merchHandler.NewHandler(services)
}

func merchandiseTypeFactory(conn *gorm.DB) merchTypeHandler.MerchandiseTypeHandler {
	data := merchTypeData.NewMySqlRecord(conn)
	services := merchTypeServices.NewService(data)
	return *merchTypeHandler.NewHandler(services)
}

func userFactory(conn *gorm.DB) userHandler.UserHandlers {
	data := userData.NewUserRepository(conn)
	services := userServices.NewService(data)
	return *userHandler.NewHandler(services)
}

func Init() handlers {
	conn := new(database.DBConf).InitDB()

	return handlers{
		MerchandiseHandler:     merchandiseFactory(conn.DB),
		MerchandiseTypeHandler: merchandiseTypeFactory(conn.DB),
		UserHandler:            userFactory(conn.DB),
	}
}
