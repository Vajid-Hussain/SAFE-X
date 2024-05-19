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
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// signupCmd represents the signup command
var signupCmd = &cobra.Command{
	Use:   "signup",
	Short: "User signup",
	Long:  `User create their accout by pass credential`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			user = reqeustmodel.User{}
		)
		reader := bufio.NewReader(os.Stdin)

		//read user name
		fmt.Printf("Enter your user name :")
		user.UserName, _ = reader.ReadString('\n')
		if user.UserName = strings.TrimSpace(user.UserName); len(user.UserName) == 0 {
			log.Fatal("user name is empty")
		}

		//read password
		fmt.Printf("Enter your password :")
		password, err := term.ReadPassword(syscall.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("")

		if user.Password = strings.TrimSpace(string(password)); len(user.Password) <= 4 {
			log.Fatal("password is less than five digit kidly strong your the password")
		}

		//read confirm password
		fmt.Printf("Re-enter your password :")
		password, err = term.ReadPassword(syscall.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("")

		if user.ConfirmPassword = strings.TrimSpace(string(string(password))); len(user.ConfirmPassword) == 0 {
			log.Fatal("confirm password is emtpy kindly enter the password")
		}

		result, err := usecase.Sighup(&user)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("signup succesfully with user id ", result)
	},
}

func init() {
	rootCmd.AddCommand(signupCmd)
}
