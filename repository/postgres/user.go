package postgres

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgconn"
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

func checkDuplicate(err error, user *domain.User) error {
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23505":
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
