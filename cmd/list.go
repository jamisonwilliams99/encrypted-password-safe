package cmd

import (
	"fmt"
	"os"

	"github.com/jamisonwilliams99/encrypted-password-safe/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all the password IDs along with what the password is used for",
	Run: func(cmd *cobra.Command, args []string) {
		passwords, err := db.AllPasswords()
		if err != nil {
			fmt.Println("Something went wrong: ", err)
			os.Exit(1)
		}

		if len(passwords) == 0 {
			fmt.Println("No passwords are currently stored in the password safe")
			return
		}

		for _, pw := range passwords {
			fmt.Printf("%d. %s\n", pw.Id, pw.UsedFor)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
