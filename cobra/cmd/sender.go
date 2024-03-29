/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// senderCmd represents the sender command
var senderCmd = &cobra.Command{
	Use:   "sender",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sender:Run")
	},

	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("sender:PreRun")
	},
	// PersistentPreRun: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("sender:PersistentPreRun")
	// },
	// PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
	// 	cmd.Context()
	// 	fmt.Println("sender:PersistentPreRunE")
	// 	return nil
	// },
}

func init() {
	rootCmd.AddCommand(senderCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// senderCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// senderCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
