package domain

import "github.com/labstack/echo/v4"

type User struct {
	ID       int
	Nama     string `validate:"requered"`
	Username string `validate:"requered"`
	Email    string `validate:"requered"`
	Password string `validate:"requered"`
	No_HP    string
}

type UserHandler interface {
	Register() echo.HandlerFunc
	GetSpecificUser() echo.HandlerFunc
}

type UserUseCases interface {
	Register(newUser User) (User, error)
	GetSpecificUser(userId int) (User, error)
	Login(email string, password string) (username string, token string, err error)
}

type UserData interface {
	Register(newUser User) (User, error)
	GetSpecificUser(userId int) (User, error)
	Login(email string, password string) (username string, token string, err error)
}
