package reqeustmodel

type User struct {
	UserName        string
	Password        string
	ConfirmPassword string
}

type Token struct {
	Token string `json:"token"`
}

type Credential struct {
	UserID     string
	Name       string
	Secret     string
	CipherText []byte
}

type GetSecret struct {
	UserID string
	Name   string
}

type GetKey struct {
	UserID string
}
