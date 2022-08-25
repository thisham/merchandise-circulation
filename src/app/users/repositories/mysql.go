package repositories

import (
	"merchandise-circulation-api/src/app/users"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

// InsertUser implements users.Repositories
func (repo *repository) InsertUser(user users.Domain) (
	userID string, err error) {
	record := mapToRecord(user)
	record.ID = uuid.Must(uuid.NewRandom())
	if err := repo.db.Create(&record).Error; err != nil {
		return "", err
	}

	return record.ID.String(), nil
}

// SelectUserByID implements users.Repositories
func (repo *repository) SelectUserByID(id string) (
	user users.Domain, err error) {
	var record User

	if err = repo.db.First(&record, id).Error; record.ID == uuid.Nil {
		return users.Domain{}, err
	}
	return record.mapToDomain(), nil
}

// SelectUserOnLogin implements users.Repositories
func (repo *repository) SelectUserOnLogin(email string) (
	domain users.Domain, err error) {
	var record User
	err = repo.db.Where("email", email).First(&record).Error
	return record.mapToDomain(), err
}

// SelectUsers implements users.Repositories
func (repo *repository) SelectUsers() (users []users.Domain, err error) {
	var records []User
	err = repo.db.Find(&records).Error
	return mapToBatchDomain(records), err
}

func NewUserRepository(db *gorm.DB) users.Repositories {
	return &repository{db}
}
