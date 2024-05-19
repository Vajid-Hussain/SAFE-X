package repository

import (
	reqeustmodel "github.com/Vajid-Hussain/SAFE-X/app/Models/reqeustModel"
	responsemodel "github.com/Vajid-Hussain/SAFE-X/app/Models/responseModel"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitRepository(DBs *gorm.DB) {
	db = DBs
}

func Signup(user *reqeustmodel.User) (*responsemodel.User, error) {
	var UserRes responsemodel.User
	query := "INSERT INTO safex_users (user_name, password) SELECT $1, $2 WHERE NOT EXISTS (SELECT 1 FROM safex_users WHERE user_name = $1) RETURNING *"
	result := db.Raw(query, user.UserName, user.Password).Scan(&UserRes)
	if result.Error != nil {
		return nil, responsemodel.ErrDataBase
	}

	if result.RowsAffected == 0 {
		return nil, responsemodel.ErrUserAlradyExist
	}
	return &UserRes, nil
}

func Login(user *reqeustmodel.User) (*responsemodel.Login, error) {
	var res responsemodel.Login
	query := "SELECT * FROM safex_users WHERE user_name=$1"
	result := db.Raw(query, user.UserName).Scan(&res)
	if result.Error != nil {
		return nil, responsemodel.ErrDataBase
	}

	if result.RowsAffected == 0 {
		return nil, responsemodel.ErrNoUserExist
	}
	return &res, nil
}

func StoreSecret(credential *reqeustmodel.Credential) error {
	// credential.Secret = "noting"
	query := "INSERT INTO safex_stores (user_ids, name, secret) SELECT $1, $2, $3 WHERE NOT EXISTS(SELECT 1 FROM safex_stores WHERE user_ids=$1 AND name=$2)"
	// query:="INSERT INTO safex_stores (user_id, name, secret) VALUES($1,$2, $3)"
	result := db.Exec(query, credential.UserID, credential.Name, credential.CipherText)
	if result.Error != nil {
		return responsemodel.ErrDataBase
	}

	if result.RowsAffected == 0 {
		return responsemodel.ErrSecretKeyRepeat
	}
	return nil
}

func FetchSecret(req *reqeustmodel.GetSecret) (*responsemodel.Secret, error) {
	var res responsemodel.Secret
	query := "SELECT * FROM safex_stores WHERE name= $1 AND user_ids = $2 "
	result := db.Raw(query, req.Name, req.UserID).Scan(&res)
	if result.Error != nil {
		return nil, responsemodel.ErrDataBase
	}

	if result.RowsAffected == 0 {
		return nil, responsemodel.ErrNoMatchingSecret
	}
	return &res, nil
}

func AllKey(req reqeustmodel.GetKey) (*responsemodel.SecretsCollecton, error) {
	var res responsemodel.SecretsCollecton
	query := "SELECT name FROM safex_stores WHERE user_ids= $1"
	result := db.Raw(query, req.UserID).Scan(&res.Name)
	if result.Error != nil {
		return nil, responsemodel.ErrDataBase
	}

	if result.RowsAffected == 0 {
		return nil, responsemodel.ErrNoSecret
	}
	return &res, nil
}

func DeleteSecret(req *reqeustmodel.GetSecret) error {
	query := "DELETE FROM safex_stores WHERE user_ids =$1 AND name=$2"
	result := db.Exec(query, req.UserID, req.Name)
	if result.Error != nil {
		return responsemodel.ErrDataBase
	}

	if result.RowsAffected == 0 {
		return responsemodel.ErrNoMatchingSecret
	}
	return nil
}
