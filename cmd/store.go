/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// storeCmd represents the store command
var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Store name and password",
	Long:  `Store name and password`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("store called")
	},
}

func init() {
	rootCmd.AddCommand(storeCmd)
}
