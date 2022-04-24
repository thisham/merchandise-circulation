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
)

type handlers struct {
	MerchandiseHandler     merchHandler.MerchandiseHandler
	MerchandiseTypeHandler merchTypeHandler.MerchandiseTypeHandler
}

func merchandiseFactory() merchHandler.MerchandiseHandler {
	data := merchData.NewMySqlRecord(database.DB)
	services := merchServices.NewService(data)
	return *merchHandler.NewHandler(services)
}

func merchandiseTypeFactory() merchTypeHandler.MerchandiseTypeHandler {
	data := merchTypeData.NewMySqlRecord(database.DB)
	services := merchTypeServices.NewService(data)
	return *merchTypeHandler.NewHandler(services)
}

func Init() handlers {
	return handlers{
		MerchandiseHandler:     merchandiseFactory(),
		MerchandiseTypeHandler: merchandiseTypeFactory(),
	}
}
