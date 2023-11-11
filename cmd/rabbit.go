/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rabbitCmd represents the rabbit command
var rabbitCmd = &cobra.Command{
	Use:   "rabbit",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rabbit called")
	},
}

func init() {
	rootCmd.AddCommand(rabbitCmd)
}
