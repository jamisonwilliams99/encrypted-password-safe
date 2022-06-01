package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates a password that is stored in the password safe",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Updated the password")
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
