package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "password",
	Short: "This is a CLI password vault",
}
