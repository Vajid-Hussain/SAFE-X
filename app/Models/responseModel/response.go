package responsemodel

import (
	"errors"
)

var (
	ErrDataBase        = errors.New("someting went wrong please try again later")
	ErrUserAlradyExist = errors.New("user name alrady exist")
)

type User struct {
	UserID   string
	UserName string
}
