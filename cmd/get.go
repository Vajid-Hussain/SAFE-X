/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	reqeustmodel "github.com/Vajid-Hussain/SAFE-X/app/Models/reqeustModel"
	responsemodel "github.com/Vajid-Hussain/SAFE-X/app/Models/responseModel"
	"github.com/Vajid-Hussain/SAFE-X/app/usecase"
	"github.com/Vajid-Hussain/SAFE-X/app/utils"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve a single secret",
	Long:  `Retrieve a single secret from the database by providing the necessary key or identifier.`,	
	Run: func(cmd *cobra.Command, args []string) {
		var (
			req reqeustmodel.GetSecret
			err error
		)

		req.UserID, err = utils.ValidateToken()
		if err != nil {
			log.Fatal(err)
		}

		reader := bufio.NewReader(os.Stdin)

		fmt.Printf("Enter name :")
		req.Name, _ = reader.ReadString('\n')
		if req.Name = strings.TrimSpace(req.Name); len(req.Name) == 0 {
			log.Fatal("Name is empty")
		}

		//get data
		result, err := usecase.GetSecret(&req)
		if errors.Is(err, responsemodel.ErrNoMatchingSecret) {
			fmt.Println(responsemodel.ErrNoMatchingSecret)
			return
		}

		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("value : %s\n", result.SecretPlainText)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
