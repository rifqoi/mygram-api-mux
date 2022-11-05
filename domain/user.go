package domain

import (
	"context"
	"time"
)

type UserRepository interface {
	InsertUser(context.Context, UserCreateParams) error
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

type UserUpdate struct {
}
