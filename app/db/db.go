package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type safex_users struct {
	UserID   int    `gorm:"type:bigint;primary_Key,AUTO_INCREMENT"`
	UserName string `gorm:"unique"`
	Password string
}

func InitDB(dbconnection string) (*gorm.DB, error) {
	DB, err := gorm.Open(postgres.Open(dbconnection), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = DB.AutoMigrate(&safex_users{})
	if err != nil {
		return nil, err
	}

	return DB, nil
}
