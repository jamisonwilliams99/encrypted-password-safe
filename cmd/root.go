package cmd

import (
	"github.com/jamisonwilliams99/encrypted-password-safe/db"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "password",
	Short: "This is a CLI password vault",
}

// used in the update and remove command to verify the password that the user enters
func verifyPw(id int, password string, key string) (bool, error) {
	decryptedPw, _, err := db.RetrievePassword(id, key) // here we don't care about the encrpyted password
	if err != nil {
		return false, err
	}

	return (password == decryptedPw), nil
}
