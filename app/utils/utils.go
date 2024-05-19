package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	reqeustmodel "github.com/Vajid-Hussain/SAFE-X/app/Models/reqeustModel"
	"github.com/Vajid-Hussain/SAFE-X/app/config"
	"github.com/absagar/go-bcrypt"
	"github.com/golang-jwt/jwt"
)

var configData config.Config

func LoadConfig(config *config.Config) {
	configData = *config
}

func CreateToken(userID, secretKey string) (string, error) {
	secretKeyByte := []byte(secretKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userID": userID,
			"exp":    time.Now().Add(time.Hour * 1).Unix(),
		})

	tokenString, err := token.SignedString(secretKeyByte)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString, secretKey string) (string, error) {
	secretKeyByte := []byte(secretKey)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKeyByte, nil
	})

	if err != nil {
		return "", fmt.Errorf("%s please login ", err)
	}

	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", err
	}

	userID, ok := claims["userID"].(string)
	if !ok {
		return "", fmt.Errorf("token claims is not matching")
	}

	return userID, nil
}

func HashPassword(password string) (string, error) {
	return bcrypt.Hash(password)
}

func MatchHashPassword(password, hastPassword string) bool {
	return bcrypt.Match(password, hastPassword)
}

func ValidateToken() (string, error) {
	var token reqeustmodel.Token

	// get user home dir
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	// read data from config
	filePath := filepath.Join(homeDir, configData.ConfigFilePath)
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("file paht", filePath)
		return "", err
	}

	//verify token
	json.Unmarshal(data, &token)

	userID, err := VerifyToken(token.Token, configData.JwtSecret)
	if err != nil {
		return "", err
	}
	return userID, nil
}

func Erncypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	b := base64.StdEncoding.EncodeToString(text)
	cypherText := make([]byte, aes.BlockSize+len(b))
	iv := cypherText[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(cypherText[aes.BlockSize:], []byte(b))
	return cypherText, nil
}

func DEcrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}
	return data, nil
}
