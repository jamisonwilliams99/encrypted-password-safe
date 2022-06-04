package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jamisonwilliams99/encrypted-password-safe/db"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes an ecrypted password from the password safe database",
	Run: func(cmd *cobra.Command, args []string) {
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

		err = db.DeletePassword(id)
		if err != nil {
			fmt.Println("Something went wrong: ", err)
			os.Exit(1)
		}
		fmt.Println("Password successfully deleted")
	},
}

func init() {
	RootCmd.AddCommand(removeCmd)
}
