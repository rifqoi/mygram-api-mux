package domain

import (
	"context"
	"time"

	"github.com/rifqoi/mygram-api-mux/repository/postgres/db"
)

type UserRepository interface {
	InsertUser(context.Context, UserCreateParams) error
	FindUserByEmail(ctx context.Context, email string) (*User, error)
	UpdateUser(ctx context.Context, userToUpdate db.UpdateUserByIDParams) (*UserUpdateResponse, error)
	FindUserByID(ctx context.Context, id int) (*User, error)
}

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type UserCreateParams struct {
	Age      int    `json:"age" validate:"required,gte=8"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,validatepassword"`
	Username string `json:"username" validate:"required"`
}

type UserCreateResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Age      int    `json:"age"`
}

type UserLoginParams struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,validatepassword"`
}

type UserUpdateParams struct {
	ID       int    `json:"id" validate:"required"`
	Age      int    `json:"age" validate:"required,gte=8"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,validatepassword"`
	Username string `json:"username" validate:"required"`
}

type UserUpdateResponse struct {
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
