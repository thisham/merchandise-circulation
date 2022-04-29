package data

import "merchandise-circulation-api/src/app/merchandise_types"

// model
type MerchandiseType struct {
	ID   int
	Name string
}

func mapToDomain(record MerchandiseType) merchandise_types.Domain {
	return merchandise_types.Domain{
		ID:   record.ID,
		Name: record.Name,
	}
}

func mapToBatchDomain(records []MerchandiseType) []merchandise_types.Domain {
	domain := []merchandise_types.Domain{}

	for _, rec := range records {
		domain = append(domain, mapToDomain(rec))
	}
	return domain
}
