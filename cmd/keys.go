/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"log"

	reqeustmodel "github.com/Vajid-Hussain/SAFE-X/app/Models/reqeustModel"
	responsemodel "github.com/Vajid-Hussain/SAFE-X/app/Models/responseModel"
	"github.com/Vajid-Hussain/SAFE-X/app/usecase"
	"github.com/Vajid-Hussain/SAFE-X/app/utils"
	"github.com/spf13/cobra"
)

// keysCmd represents the keys command
var keysCmd = &cobra.Command{
	Use:   "keys",
	Short: "List all stored keys",
	Long:  `Retrieve and display all keys stored in the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			req reqeustmodel.GetKey
			err error
		)

		//validate user token
		req.UserID, err = utils.ValidateToken()
		if err != nil {
			log.Fatal(err)
		}

		// fetch all keys
		result, err := usecase.AllKey(req)
		if errors.Is(err, responsemodel.ErrNoSecret) {
			fmt.Println(responsemodel.ErrNoSecret)
			return
		}

		if err != nil {
			log.Fatal(err)
		}

		for _, val := range result.Name {
			fmt.Printf("key :%s\n", val)
		}
	},
}

func init() {
	rootCmd.AddCommand(keysCmd)
}
