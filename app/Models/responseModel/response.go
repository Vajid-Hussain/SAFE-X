package responsemodel

import (
	"errors"
)

var (
	ErrDataBase         = errors.New("someting went wrong please try again later")
	ErrUserAlradyExist  = errors.New("user name alrady exist")
	ErrNoUserExist      = errors.New("no user exist")
	ErrPasswordNotMatch = errors.New("password not match")
)

type User struct {
	UserID   string
	UserName string
}

type Login struct{
	UserID string
	UserName string
	Password string
}