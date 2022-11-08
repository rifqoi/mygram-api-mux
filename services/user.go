package services

import (
	"context"
	"database/sql"
	"errors"

	"github.com/rifqoi/mygram-api-mux/domain"
	"github.com/rifqoi/mygram-api-mux/helpers"
	"github.com/rifqoi/mygram-api-mux/repository/postgres/db"
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

func (u *UserService) UpdateUser(ctx context.Context, currentUserID int, req domain.UserUpdateParams) (*domain.UserUpdateResponse, error) {
	userToUpdate := db.UpdateUserByIDParams{
		ID: int32(currentUserID),
		Email: sql.NullString{
			String: req.Email,
			Valid:  checkNullType(req.Email),
		},
		Password: sql.NullString{
			String: req.Password,
			Valid:  checkNullType(req.Password),
		},
		Username: sql.NullString{
			String: req.Username,
			Valid:  checkNullType(req.Username),
		},
		Age: sql.NullInt32{
			Int32: int32(req.Age),
			Valid: checkNullType(req.Age),
		},
	}
	updatedUser, err := u.repo.UpdateUser(ctx, userToUpdate)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (u *UserService) DeleteUser(ctx context.Context, id int) error {
	err := u.repo.DeleteUserByID(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

// https://stackoverflow.com/questions/71186757/how-to-check-if-the-value-of-a-generic-type-is-the-zero-value
func checkNullType[T comparable](v T) bool {
	// new(T) will return a pointer type of type T
	// then we will compare the value of v and dereferenced T (*new(t))
	// dereferenced T will return the nil type of T (0 for int, "" for string)
	// so if v is a nil value of the type, it will return true
	t := *new(T)
	return v != t
}
