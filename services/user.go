package services

import (
	"context"

	"github.com/rifqoi/mygram-api-mux/domain"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) InsertUser(ctx context.Context, req domain.UserCreateParams) (*domain.UserCreateResponse, error) {
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
