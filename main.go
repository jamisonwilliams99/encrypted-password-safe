package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jamisonwilliams99/encrypted-password-safe/cmd"
	"github.com/jamisonwilliams99/encrypted-password-safe/db"
	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "passwords.db")
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
