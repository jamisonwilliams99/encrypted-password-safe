package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates a password that is stored in the password safe",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Updated the password")
		if len(args) < 3 {
			fmt.Println("the add command requires 3 arguments: id, password, key")
			os.Exit(1)
		}
		id, _ := strconv.Atoi(args[0])
		password, key := args[1], args[2]

		isCorrectPw, err := verifyPw(id, password, key)
		if err != nil {
			fmt.Println("Something went wrong:", err)
			os.Exit(1)
		}
		if !isCorrectPw {
			fmt.Println("The password or key that you entered is incorrect")
			os.Exit(1)
		}

		// Finish this command

	},
}

func init() {
	RootCmd.AddCommand(updateCmd)
}
