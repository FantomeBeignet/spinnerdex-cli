/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"spinnerdex-cli/utils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var board string
var twitter string
var youtube string

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
		if len(args) == 0 {
			color.Red("Please provide a name for the spinner")
		} else if len(args) >= 2 {
			color.Red("Please provide only one name for the spinner")
		} else {
			name := args[0]
			board := board
			twitter := twitter
			youtube := youtube
			apiKey := viper.GetString("api-key")
			spinner := utils.Spinner{
				Name:    name,
				Twitter: twitter,
				Youtube: youtube,
				Board:   board,
			}
			utils.AddSpinner(spinner, apiKey)
		}
	},
}

func init() {
	spinnerCmd.AddCommand(spinnerAddCmd)

	spinnerAddCmd.Flags().StringVarP(&board, "board", "b", "", "The board the spinner is from")
	spinnerAddCmd.MarkFlagRequired("board")

	spinnerAddCmd.Flags().StringVarP(&twitter, "twitter", "t", "", "The Twitter link of the spinner")
	spinnerAddCmd.Flags().StringVarP(&youtube, "youtube", "y", "", "The YouTube link of the spinner")
}
