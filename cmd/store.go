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
	"syscall"

	reqeustmodel "github.com/Vajid-Hussain/SAFE-X/app/Models/reqeustModel"
	"github.com/Vajid-Hussain/SAFE-X/app/usecase"
	"github.com/Vajid-Hussain/SAFE-X/app/utils"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// storeCmd represents the store command
var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Save a name and secret",
	Long:  `Store a name and secret securely in the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			credential reqeustmodel.Credential
			err        error
		)

		credential.UserID, err = utils.ValidateToken()
		if err != nil {
			log.Fatal(err)
		}

		reader := bufio.NewReader(os.Stdin)

		fmt.Printf("Enter name :")
		credential.Name, _ = reader.ReadString('\n')
		if credential.Name = strings.TrimSpace(credential.Name); len(credential.Name) == 0 {
			log.Fatal("user name is empty")
		}

		//read password
		fmt.Printf("Enter secret :")
		password, err := term.ReadPassword(syscall.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("")

		if credential.Secret = strings.TrimSpace(string(password)); len(credential.Secret) <= 1 {
			log.Fatal("secret is empty")
		}

		err = usecase.StoreSecret(&credential)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s stored succesfully \n", credential.Name)
	},
}

func init() {
	rootCmd.AddCommand(storeCmd)
}
