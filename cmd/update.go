package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jamisonwilliams99/encrypted-password-safe/db"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates a password that is stored in the password safe",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 4 {
			fmt.Println("the add command requires 4 arguments: id, password, new password, key")
			os.Exit(1)
		}
		id, _ := strconv.Atoi(args[0])
		password, newPassword, key := args[1], args[2], args[3]

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
		err = db.UpdatePassword(id, newPassword, key)
		if err != nil {
			fmt.Println("Something went wrong:", err)
			os.Exit(1)
		}
		fmt.Println("Password updated successfully")
	},
}

func init() {
	RootCmd.AddCommand(updateCmd)
}
