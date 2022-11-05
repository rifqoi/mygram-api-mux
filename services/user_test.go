package services

import (
	"context"
	"errors"
	"testing"

	"github.com/rifqoi/mygram-api-mux/domain"
	"github.com/rifqoi/mygram-api-mux/domain/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserTestSuite struct {
	suite.Suite
	mockRepo    *mocks.UserRepository
	UserService *UserService
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, &UserTestSuite{})
}

func (ts *UserTestSuite) SetupTest() {
	mockRepo := new(mocks.UserRepository)
	userService := NewUserService(mockRepo)

	ts.mockRepo = mockRepo
	ts.UserService = userService

}

func (ts *UserTestSuite) TestInsertUser() {
	createParams := domain.UserCreateParams{
		Username: "rifqoi",
		Email:    "rifqoi@gmail.com",
		Password: "Rifqoi@123",
		Age:      21,
	}
	expectedCreateResponse := &domain.UserCreateResponse{
		Email:    "rifqoi@gmail.com",
		Username: "rifqoi",
		Age:      21,
	}

	ts.Run("InsertUser_Success", func() {
		mockRepo := new(mocks.UserRepository)
		userService := NewUserService(mockRepo)

		mockRepo.On("InsertUser", mock.Anything, mock.AnythingOfType("domain.UserCreateParams")).Return(nil)

		actualResp, err := userService.InsertUser(context.Background(), createParams)
		ts.Equal(expectedCreateResponse, actualResp)
		ts.Nil(err)
	})

	ts.Run("InsertUser_DuplicateEmail", func() {
		mockRepo := new(mocks.UserRepository)
		userService := NewUserService(mockRepo)
		errorDuplicateEmail := errors.New("User with email rifqoi@gmail.com already exists.")

		mockRepo.On("InsertUser", mock.Anything, mock.AnythingOfType("domain.UserCreateParams")).Return(errorDuplicateEmail)

		actualResp, err := userService.InsertUser(context.Background(), createParams)

		ts.NotNil(err)
		ts.Nil(actualResp)
		ts.Equal(errorDuplicateEmail, err)
	})

	ts.Run("InsertUser_DuplicateUsername", func() {
		mockRepo := new(mocks.UserRepository)
		userService := NewUserService(mockRepo)
		errorDuplicateUsername := errors.New("User with username rifqoi already exists.")

		mockRepo.On("InsertUser", mock.Anything, mock.AnythingOfType("domain.UserCreateParams")).Return(errorDuplicateUsername)

		actualResp, err := userService.InsertUser(context.Background(), createParams)

		ts.NotNil(err)
		ts.Nil(actualResp)
		ts.Equal(errorDuplicateUsername, err)
	})

	ts.Run("InsertUser_ErrorGeneral", func() {
		mockRepo := new(mocks.UserRepository)
		userService := NewUserService(mockRepo)

		mockRepo.On("InsertUser", mock.Anything, mock.AnythingOfType("domain.UserCreateParams")).Return(errors.New("error"))

		resp, err := userService.InsertUser(context.Background(), createParams)
		ts.Nil(resp)
		ts.NotNil(err)
	})
}
