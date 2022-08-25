package services

import (
	"errors"
	"merchandise-circulation-api/src/app/users"
	"merchandise-circulation-api/src/app/users/mocks"
	"merchandise-circulation-api/src/utils"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	mockRepo                 mocks.Repositories
	services                 users.Services
	sampleDomain             users.Domain
	sampleDomainList         []users.Domain
	sampleDomainUnregistered users.Domain
	sampleUUID               uuid.UUID
	sampleEmail              string
	samplePassword           string
)

func TestMain(m *testing.M) {
	services = NewService(&mockRepo)
	sampleUUID = uuid.Must(uuid.NewRandom())
	sampleEmail = "user@example.com"
	samplePassword = "thestrongestpassword"

	sampleDomain = users.Domain{
		ID:        sampleUUID,
		Name:      "Sample User",
		Email:     sampleEmail,
		Password:  samplePassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	sampleDomainUnregistered = users.Domain{
		Name:      "Sample User Two",
		Email:     "user.two@example.com",
		Password:  "thestrongestpassword",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	sampleDomainList = []users.Domain{sampleDomain, {
		ID:        uuid.Must(uuid.NewRandom()),
		Name:      sampleDomainUnregistered.Name,
		Email:     sampleDomainUnregistered.Email,
		Password:  sampleDomainUnregistered.Password,
		CreatedAt: sampleDomainUnregistered.CreatedAt,
		UpdatedAt: sampleDomainUnregistered.UpdatedAt,
	}}

	os.Exit(m.Run())
}

func TestAttemptLogin(t *testing.T) {
	t.Run("should got auth token", func(t *testing.T) {
		hashedPassphrase, _ := utils.HashMake(samplePassword)
		mockRepo.On("SelectUserOnLogin", sampleEmail).Return(hashedPassphrase, nil).
			Once()
		token, err := services.AttemptLogin(sampleEmail, samplePassword)

		assert.NotEqual(t, "", token)
		assert.Nil(t, err)
	})

	t.Run("should got incorrect passphrase", func(t *testing.T) {
		hashedPassphrase, _ := utils.HashMake(samplePassword)
		mockRepo.On("SelectUserOnLogin", sampleEmail).Return(hashedPassphrase, nil).
			Once()
		token, err := services.AttemptLogin(sampleEmail, "anotherpasshrase")

		assert.Equal(t, "", token)
		assert.NotNil(t, err)
	})

	t.Run("should got user not found", func(t *testing.T) {
		mockRepo.On("SelectUserOnLogin", sampleEmail).
			Return("", errors.New("record not found")).Once()
		token, err := services.AttemptLogin(sampleEmail, "anotherpasshrase")

		assert.Equal(t, "", token)
		assert.NotNil(t, err)
	})
}

func TestAttemptRegister(t *testing.T) {
	t.Run("should got registered", func(t *testing.T) {
		mockRepo.On("SelectUserOnLogin", sampleDomainUnregistered.Email).
			Return("", errors.New("record not found")).Once()
		mockRepo.On("InsertUser", sampleDomainUnregistered).Return(sampleUUID, nil).
			Once()
		token, err := services.AttemptRegister(sampleDomainUnregistered)

		assert.NotEqual(t, "", token)
		assert.Nil(t, err)
	})

	t.Run("should got registered user error", func(t *testing.T) {
		hashedPassphrase, _ := utils.HashMake(samplePassword)
		mockRepo.On("SelectUserOnLogin", sampleDomainUnregistered.Email).
			Return(hashedPassphrase, nil).Once()
		mockRepo.On("InsertUser", sampleDomainUnregistered).
			Return("", errors.New("email was used")).Once()
		token, err := services.AttemptRegister(sampleDomainUnregistered)

		assert.Equal(t, "", token)
		assert.NotNil(t, err)
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("SelectUserOnLogin", sampleDomainUnregistered.Email).
			Return("", errors.New("database error")).Once()
		token, err := services.AttemptRegister(sampleDomainUnregistered)

		assert.Equal(t, "", token)
		assert.NotNil(t, err)
	})
}

func TestGetAllUsers(t *testing.T) {
	t.Run("should got all users", func(t *testing.T) {
		mockRepo.On("SelectUsers").Return(sampleDomainList).Once()
		domains, err := services.GetAllUsers()

		assert.NotNil(t, domains)
		assert.Nil(t, err)
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("SelectUsers").Return(sampleDomainList).Once()
		domain, err := services.GetAllUsers()

		assert.Nil(t, domain)
		assert.NotNil(t, err)
	})
}

func TestGetUserByID(t *testing.T) {
	t.Run("should got selected user", func(t *testing.T) {
		mockRepo.On("SelectUserByID", sampleUUID).Return(sampleDomain, nil).Once()
		domain, err := services.GetUserByID(sampleUUID.String())

		assert.NotNil(t, domain)
		assert.Nil(t, err)
	})

	t.Run("should got user not found error", func(t *testing.T) {
		mockRepo.On("SelectUserByID", sampleUUID).
			Return(users.Domain{}, errors.New("record not found")).Once()
		domain, err := services.GetUserByID(sampleUUID.String())

		assert.NotNil(t, domain)
		assert.Nil(t, err)
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("SelectUserByID", sampleUUID).
			Return(users.Domain{}, errors.New("database error")).Once()
		domain, err := services.GetUserByID(sampleUUID.String())

		assert.Nil(t, domain)
		assert.NotNil(t, err)
	})
}
