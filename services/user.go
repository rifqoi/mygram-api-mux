package services

import (
	"context"
	"errors"

	"github.com/rifqoi/mygram-api-mux/domain"
	"github.com/rifqoi/mygram-api-mux/helpers"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) FindUserByID(ctx context.Context, id int) (*domain.User, error) {
	user, err := u.repo.FindUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserService) RegisterUser(ctx context.Context, req domain.UserCreateParams) (*domain.UserCreateResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	req.Password = string(hashedPassword)

	err = u.repo.InsertUser(ctx, req)
	if err != nil {
		return nil, err
	}

	res := &domain.UserCreateResponse{
		Email:    req.Email,
		Username: req.Username,
		Age:      req.Age,
	}

	return res, nil
}

func (u *UserService) Login(ctx context.Context, req domain.UserLoginParams) (*string, error) {
	user, err := u.repo.FindUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("Password doesn't match!")
	}

	token, err := helpers.GenerateToken(user.Email, user.ID)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

// func (u *UserService) UpdateUser(ctx context.Context, req domain.UserUpdateParams) (*domain.User, error) {
// 	userToUpdate := db.UpdateUserByIDParams{
// 		ID: int32(req.ID),
// 	}
// 	updatedUser, err := u.repo.UpdateUser(ctx)
// }
