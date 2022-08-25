package repositories

import (
	"merchandise-circulation-api/src/app/users"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID
	Name     string
	Email    string `gorm:"unique;not null;type:varchar(100)"`
	Password string
}

func mapToRecord(domain users.Domain) User {
	return User{
		ID:       domain.ID,
		Name:     domain.Name,
		Email:    domain.Email,
		Password: domain.Password,
	}
}

func (rec *User) mapToDomain() users.Domain {
	return users.Domain{
		ID:        rec.ID,
		Name:      rec.Name,
		Email:     rec.Email,
		Password:  rec.Password,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func mapToBatchDomain(records []User) []users.Domain {
	var domains []users.Domain

	for _, record := range records {
		domains = append(domains, record.mapToDomain())
	}
	return domains
}
