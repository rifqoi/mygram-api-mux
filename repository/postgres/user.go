package postgres

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/rifqoi/mygram-api-mux/domain"
	"github.com/rifqoi/mygram-api-mux/repository/postgres/db"
)

type userRepository struct {
	query *db.Queries
}

func NewUserRepository(query *db.Queries) domain.UserRepository {
	return &userRepository{
		query: query,
	}
}

func (u *userRepository) InsertUser(ctx context.Context, req domain.UserCreateParams) error {
	userToInsert := db.InsertUserParams{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Age:      int32(req.Age),
	}

	err := u.query.InsertUser(ctx, userToInsert)
	if err != nil {
		err = checkDuplicate(err, &domain.User{
			Email:    req.Email,
			Username: req.Username,
		})
		return err
	}

	return nil
}

func (u *userRepository) FindUserByID(ctx context.Context, id int) (*domain.User, error) {
	res, err := u.query.FindUserById(ctx, int32(id))
	if err != nil {
		return nil, errors.New("User not found!")
	}

	user := &domain.User{
		ID:       int(res.ID),
		Email:    res.Email,
		Username: res.Username,
		Password: res.Password,
		Age:      int(res.Age),
	}

	return user, nil
}
func (u *userRepository) FindUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	res, err := u.query.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("User not found!")
	}

	user := &domain.User{
		ID:       int(res.ID),
		Email:    res.Email,
		Username: res.Username,
		Password: res.Password,
		Age:      int(res.Age),
	}

	return user, nil
}

func (u *userRepository) UpdateUser(ctx context.Context, userToUpdate db.UpdateUserByIDParams) (*domain.UserUpdateResponse, error) {
	updatedUser, err := u.query.UpdateUserByID(ctx, userToUpdate)
	if err != nil {
		return nil, err
	}
	resp := &domain.UserUpdateResponse{
		Email:    updatedUser.Email,
		Username: updatedUser.Username,
		Age:      int(updatedUser.Age),
	}
	return resp, nil
}

func (u *userRepository) DeleteUserByID(ctx context.Context, id int) error {
	err := u.query.DeleteUserByID(ctx, int32(id))
	if err != nil {
		return err
	}
	return nil
}

func checkDuplicate(err error, user *domain.User) error {
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case pgerrcode.UniqueViolation:
				if strings.Contains(pgErr.Message, "username") {
					err = fmt.Errorf("User with username %s already exists.", user.Username)
				} else if strings.Contains(pgErr.Message, "email") {
					err = fmt.Errorf("User with email %s already exists.", user.Email)
				}
				return err
			}
		}
		return err
	}
	return err
}
