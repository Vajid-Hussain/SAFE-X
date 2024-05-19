package usecase

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"

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

	user.UserName = strings.ToLower(user.UserName)

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

	// create jwt token
	jwtToken, err := utils.CreateToken(result.UserID, configData.JwtSecret)
	if err != nil {
		return nil, err
	}

	tokenModel := reqeustmodel.Token{Token: jwtToken}
	byteTokenModel, err := json.Marshal(tokenModel)
	if err != nil {
		return nil, err
	}

	// create config file store token
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
		return nil, err
	}

	_, err = file.Write(byteTokenModel)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func Login(user *reqeustmodel.User) error {
	user.UserName = strings.ToLower(user.UserName)

	userData, err := repository.Login(user)
	if err != nil {
		return err
	}

	match := utils.MatchHashPassword(user.Password, userData.Password)
	if !match {
		return responsemodel.ErrPasswordNotMatch
	}

	//create jwt token
	jwtToken, err := utils.CreateToken(userData.UserID, configData.JwtSecret)
	if err != nil {
		return err
	}

	tokenModel := reqeustmodel.Token{Token: jwtToken}
	byteTokenModel, err := json.Marshal(tokenModel)
	if err != nil {
		return err
	}

	// store in config file
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configDir := filepath.Join(userHomeDir, configData.ConfigPath)
	confFileDir := filepath.Join(userHomeDir, configData.ConfigFilePath)

	err = os.MkdirAll(configDir, os.ModePerm)
	if err != nil {
		return err
	}

	file, err := os.Create(confFileDir)
	if err != nil {
		return err
	}

	_, err = file.Write(byteTokenModel)
	if err != nil {
		return err
	}

	return nil
}
