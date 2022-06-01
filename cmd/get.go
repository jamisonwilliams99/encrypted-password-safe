package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jamisonwilliams99/encrypted-password-safe/db"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieves and decrypts an encrypted password from the password safe",
	Run: func(cmd *cobra.Command, args []string) {
		showDecryption, _ := cmd.Flags().GetBool("show")

		if len(args) < 2 {
			fmt.Println("the get command requires 2 arguments: id, key")
			os.Exit(1)
		}

		id, _ := strconv.Atoi(args[0])
		key := args[1]
		password, encryptedPassword, err := db.RetrievePassword(id, key)
		if err != nil {
			fmt.Println("Something went wrong: ", err)
			os.Exit(1)
		}

		if showDecryption {
			fmt.Printf("\n%s   ->   %s\n\n", encryptedPassword, password)
		}

		fmt.Printf("requested password: %s\n", password)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
	getCmd.PersistentFlags().Bool("show", false, "used for user to show the password decryption")
}
