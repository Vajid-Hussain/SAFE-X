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
