/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	reqeustmodel "github.com/Vajid-Hussain/SAFE-X/app/Models/reqeustModel"
	"github.com/Vajid-Hussain/SAFE-X/app/usecase"
	"github.com/Vajid-Hussain/SAFE-X/app/utils"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long:  `A longn.`,
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
			log.Fatal("user name is empty")
		}

		result, err := usecase.GetSecret(&req)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("value : %s\n", result.SecretPlainText)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
