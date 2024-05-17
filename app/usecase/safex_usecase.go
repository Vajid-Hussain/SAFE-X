package usecase

import (
	"fmt"
	"log"

	reqeustmodel "github.com/Vajid-Hussain/SAFE-X/app/Models/reqeustModel"
	responsemodel "github.com/Vajid-Hussain/SAFE-X/app/Models/responseModel"
	"github.com/Vajid-Hussain/SAFE-X/app/config"
	"github.com/Vajid-Hussain/SAFE-X/app/repository"
	"github.com/Vajid-Hussain/SAFE-X/app/utils"
)

var configData *config.Config

func InitConfig(credential *config.Config) {
	configData = credential
}

func Sighup(user *reqeustmodel.User) (*responsemodel.User, error) {
	var err error

	if user.Password != user.ConfirmPassword {
		log.Fatal("password and confirm password not matching")
	}

	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	result, err := repository.Signup(user)
	if err != nil {
		return nil, err
	}

	jwtToken, err := utils.CreateToken(result.UserID, configData.JwtSecret)
	if err != nil {
		return nil, err
	}

	fmt.Println("== token ", jwtToken)
	return result, nil
}
