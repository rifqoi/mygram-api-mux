package services

import (
	"context"
	"errors"
	"testing"

	"github.com/rifqoi/mygram-api-mux/domain"
	"github.com/rifqoi/mygram-api-mux/domain/mocks"
	"github.com/rifqoi/mygram-api-mux/helpers"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
)

type UserTestSuite struct {
	suite.Suite
	mockRepo    *mocks.UserRepository
	userService *UserService
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, &UserTestSuite{})
}

func (ts *UserTestSuite) SetupTest() {
	mockRepo := new(mocks.UserRepository)
	userService := NewUserService(mockRepo)

	ts.mockRepo = mockRepo
	ts.userService = userService
}

var (
	createParams = domain.UserCreateParams{
		Username: "rifqoi",
		Email:    "rifqoi@gmail.com",
		Password: "Rifqoi@123",
		Age:      21,
	}

	expectedCreateResponse = &domain.UserCreateResponse{
		Email:    "rifqoi@gmail.com",
		Username: "rifqoi",
		Age:      21,
	}
	loginParams = domain.UserLoginParams{
		Email:    "rifqoi@gmail.com",
		Password: "Rifqoi@123",
	}
)

func (ts *UserTestSuite) TestNewUserService() {
	mockRepo := new(mocks.UserRepository)
	userService := NewUserService(mockRepo)

	ts.IsType(&UserService{}, userService)
}

func (ts *UserTestSuite) TestRegisterUser_Success() {

	ts.mockRepo.On("InsertUser", mock.Anything, mock.AnythingOfType("domain.UserCreateParams")).Return(nil)

	actualResp, err := ts.userService.RegisterUser(context.Background(), createParams)
	ts.Equal(expectedCreateResponse, actualResp)
	ts.Nil(err)

}
func (ts *UserTestSuite) TestRegisterUser_DuplicateEmail() {

	errorDuplicateEmail := errors.New("User with email rifqoi@gmail.com already exists.")

	ts.mockRepo.On("InsertUser", mock.Anything, mock.AnythingOfType("domain.UserCreateParams")).Return(errorDuplicateEmail)

	actualResp, err := ts.userService.RegisterUser(context.Background(), createParams)

	ts.NotNil(err)
	ts.Nil(actualResp)
	ts.Equal(errorDuplicateEmail, err)
}

func (ts *UserTestSuite) TestRegisterUser_DuplicateUsername() {

	errorDuplicateUsername := errors.New("User with username rifqoi already exists.")

	ts.mockRepo.On("InsertUser", mock.Anything, mock.AnythingOfType("domain.UserCreateParams")).Return(errorDuplicateUsername)

	actualResp, err := ts.userService.RegisterUser(context.Background(), createParams)
	ts.NotNil(err)
	ts.Nil(actualResp)
	ts.Equal(errorDuplicateUsername, err)
}

func (ts *UserTestSuite) TestRegisterUser_ErrorGeneral() {

	ts.mockRepo.On("InsertUser", mock.Anything, mock.AnythingOfType("domain.UserCreateParams")).Return(errors.New("error"))

	resp, err := ts.userService.RegisterUser(context.Background(), createParams)
	ts.Nil(resp)
	ts.NotNil(err)
}

func (ts *UserTestSuite) TestLogin_Success() {
	expectedUser := &domain.User{
		ID:       1,
		Email:    "rifqoi@gmail.com",
		Username: "rifqoi",
		Age:      21,
	}
	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(loginParams.Password), bcrypt.DefaultCost)
	expectedUser.Password = string(hashedPass)
	expectedToken, _ := helpers.GenerateToken(expectedUser.Email, expectedUser.ID)

	ts.mockRepo.On("FindUserByEmail", mock.Anything, loginParams.Email).Return(expectedUser, nil)

	actualToken, err := ts.userService.Login(context.Background(), loginParams)

	ts.Nil(err)
	ts.NotNil(expectedToken)

	// actualToken returnin address of string
	// so we have to dereference it
	ts.Equal(expectedToken, *actualToken)
}

func (ts *UserTestSuite) TestLogin_UserNotFound() {
	expectedErr := errors.New("User not found!")
	ts.mockRepo.On("FindUserByEmail", mock.Anything, loginParams.Email).Return(nil, expectedErr)

	actualToken, err := ts.userService.Login(context.Background(), loginParams)

	ts.Nil(actualToken)
	ts.NotNil(err)
	ts.Equal(expectedErr, err)
	ts.EqualValues(expectedErr, err)
}

func (ts *UserTestSuite) TestLogin_PasswordNotMatch() {
	expectedErr := errors.New("Password doesn't match!")
	ts.mockRepo.On("FindUserByEmail", mock.Anything, loginParams.Email).Return(nil, expectedErr)

	actualToken, err := ts.userService.Login(context.Background(), loginParams)

	ts.Nil(actualToken)
	ts.NotNil(err)
	ts.Equal(expectedErr, err)
	ts.EqualValues(expectedErr, err)
}

func (ts *UserTestSuite) TestLogin_Token() {
	expectedErr := errors.New("error generate token")
	ts.mockRepo.On("FindUserByEmail", mock.Anything, loginParams.Email).Return(nil, expectedErr)

	actualToken, err := ts.userService.Login(context.Background(), loginParams)

	ts.Nil(actualToken)
	ts.NotNil(err)
	ts.EqualValues(expectedErr, err)
}
