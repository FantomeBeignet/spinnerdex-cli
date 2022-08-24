/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// spinnerAddCmd represents the spinnerAdd command
var spinnerAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a spinner to the database",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("spinnerAdd called")
	},
}

func init() {
	spinnerCmd.AddCommand(spinnerAddCmd)

	spinnerAddCmd.Flags().StringP("name", "n", "", "The name of the spinner")
	spinnerAddCmd.MarkFlagRequired("name")

	spinnerAddCmd.Flags().StringP("board", "b", "", "The board the spinner is from")
	spinnerAddCmd.MarkFlagRequired("board")

	spinnerAddCmd.Flags().StringP("twitter", "t", "", "The Twitter link of the spinner")
	spinnerAddCmd.Flags().StringP("youtube", "y", "", "The YouTube link of the spinner")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// spinnerAddCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// spinnerAddCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}