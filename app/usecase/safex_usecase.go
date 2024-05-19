package usecase

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

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

	tokenModel := reqeustmodel.Token{Token: jwtToken}
	byteTokenModel, err := json.Marshal(tokenModel)
	if err != nil {
		return nil, err
	}

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configDir := filepath.Join(userHomeDir, configData.ConfigPath)
	confFileDir := filepath.Join(userHomeDir, configData.ConfigFilePath)

	err = os.MkdirAll(configDir, os.ModePerm)
	if err != nil {
		return nil, err
	}
	
	file, err := os.Create(confFileDir)
	if err != nil {
	fmt.Println("----",err, confFileDir)
		return nil, err
	}

	_, err = file.Write(byteTokenModel)
	if err != nil {
		return nil, err
	}

	return result, nil
}
