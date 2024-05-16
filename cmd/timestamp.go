/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
)

// timestampCmd represents the timestamp command
var timestampCmd = &cobra.Command{
	Use:   "timestamp",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usagate the needed files
to quickly create a Cobra application.`,

	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		timezone := args[0]
		dateFlag, _ := cmd.Flags().GetString("date")
		location, err := time.LoadLocation(timezone)
		if err != nil {
			log.Fatalln("The timezone string is invalid")
		}

		var date string
		if dateFlag != "" {
			fmt.Println("worded flag")
			date = time.Now().In(location).Format(dateFlag)
		} else {
			date = time.Now().In(location).Format(time.RFC3339)[:10]
		}
		fmt.Printf("Current date in %v: %v\n", timezone, date)

	},
}

func init() {
	rootCmd.AddCommand(timestampCmd)
	timestampCmd.PersistentFlags().String("date", "", "used for specify the data")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// timestampCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// timestampCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
