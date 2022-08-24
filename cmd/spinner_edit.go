/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// spinnerEditCmd represents the spinnerEdit command
var spinnerEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a spinner in the database",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("spinnerEdit called")
	},
}

func init() {
	spinnerCmd.AddCommand(spinnerEditCmd)

	spinnerEditCmd.Flags().StringP("name", "n", "", "The name of the spinner")
	spinnerEditCmd.MarkFlagRequired("name")

	spinnerEditCmd.Flags().StringP("board", "b", "", "The board the spinner is from")
	spinnerEditCmd.Flags().StringP("twitter", "t", "", "The Twitter link of the spinner")
	spinnerEditCmd.Flags().StringP("youtube", "y", "", "The YouTube link of the spinner")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// spinnerEditCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// spinnerEditCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
