package responsemodel

import (
	"errors"
)

var (
	ErrDataBase         = errors.New("someting went wrong please try again later")
	ErrUserAlradyExist  = errors.New("the username already exists")
	ErrNoUserExist      = errors.New("no user found with the provided details")
	ErrPasswordNotMatch = errors.New("the provided password does not match our records")
	ErrSecretKeyRepeat  = errors.New("a secret with the same key already exists")
	ErrNoSecret         = errors.New("no secret found for the specified criteria")
	ErrNoMatchingSecret = errors.New("no matching secret found for the provided key")
)

type User struct {
	UserID   string
	UserName string
}

type Login struct {
	UserID   string
	UserName string
	Password string
}

type SecretsCollecton struct {
	Name []string
}

type Secret struct {
	Name            string
	Secret          []byte
	SecretPlainText string
}
