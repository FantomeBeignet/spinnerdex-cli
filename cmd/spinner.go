/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// spinnerCmd represents the spinner command
var spinnerCmd = &cobra.Command{
	Use:   "spinner",
	Short: "Command to manage spinner data",
	Long: `This command requires an API key to be set in the config file or via the --api-key flag.
This will allow you to add or edit spinners in the SpinnerDex database.`,
}

func init() {
	rootCmd.AddCommand(spinnerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// spinnerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// spinnerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
