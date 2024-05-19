/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"

	"github.com/Vajid-Hussain/SAFE-X/app/config"
	"github.com/Vajid-Hussain/SAFE-X/app/db"
	"github.com/Vajid-Hussain/SAFE-X/app/repository"
	"github.com/Vajid-Hussain/SAFE-X/app/usecase"
	"github.com/Vajid-Hussain/SAFE-X/app/utils"
	"github.com/Vajid-Hussain/SAFE-X/cmd"
)

func main() {
	config, err := config.InitConfig()
	if err != nil {
		log.Fatal("error from config ", err)
	}

	DB, err := db.InitDB(config.SupaBaseConnection)
	if err != nil {
		log.Fatal("error during connecting to database")
	}

	utils.LoadConfig(config)
	usecase.InitConfig(config)
	repository.InitRepository(DB)

	cmd.Execute()
}
