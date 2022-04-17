package factories

import (
	"merchandise-circulation-api/src/app/merchandises/data"
	"merchandise-circulation-api/src/app/merchandises/handlers"
	"merchandise-circulation-api/src/app/merchandises/services"
	"merchandise-circulation-api/src/database"
)

type Presenter struct {
	MerchandiseHandler handlers.MerchandiseHandler
}

func Init() Presenter {
	merchandiseData := data.NewMySqlRecord(database.DB)
	merchandiseService := services.NewMerchandiseServices(merchandiseData)

	return Presenter{
		MerchandiseHandler: *handlers.NewMerchandiseHandler(merchandiseService),
	}
}
