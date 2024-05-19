package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type safex_users struct {
	UserID   int    `gorm:"type:bigint;primary_Key" sql:"AUTO_INCREMENT"`
	UserName string `gorm:"unique"`
	Password string
}

type safex_store struct {
	ID     int
	UserIDs   int
	FKUser safex_users `gorm:"foreignKey:UserIDs;references:UserID"`
	Name   string
	Secret []byte
}

func InitDB(dbconnection string) (*gorm.DB, error) {
	DB, err := gorm.Open(postgres.Open(dbconnection), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = DB.AutoMigrate(&safex_users{}, safex_store{})
	if err != nil {
		return nil, err
	}

	return DB, nil
}
