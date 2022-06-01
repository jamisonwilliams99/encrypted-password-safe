package cmd

import (
	"fmt"
	"os"

	"github.com/jamisonwilliams99/encrypted-password-safe/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds an encrypted password to the password safe database",
	Run: func(cmd *cobra.Command, args []string) {
		showEncryption, _ := cmd.Flags().GetBool("show")

		if len(args) < 3 {
			fmt.Println("the add command requires 3 arguments: password, key, use")
			os.Exit(1)
		}
		password, key, usedFor := args[0], args[1], args[2]

		must(password, key)

		encryptedPassword, err := db.CreatePassword(password, key, usedFor)
		if err != nil {
			fmt.Println("Something went wrong: ", err)
			os.Exit(1)
		}

		if showEncryption {
			fmt.Printf("\n%s   ->   %s\n\n", password, encryptedPassword)
		}

		fmt.Println("Added a password to the password safe")
	},
}

func must(password string, key string) {
	if len(password) != 16 {
		fmt.Println("Password must be 16 characters")
		os.Exit(1)
	}
	if len(key) != 8 {
		fmt.Println("Key must be 8 characters")
		os.Exit(1)
	}
}

func init() {
	RootCmd.AddCommand(addCmd)
	addCmd.PersistentFlags().Bool("show", false, "used for user to show the password encryption")
}
