package cmd

import (
	"github.com/spf13/cobra"
	"tctl/sheety"
)
var createCmd = &cobra.Command{
	Use:   "create",
	Run: func(cmd *cobra.Command, args []string) { 
		sheety.CreateTask("Work on helm-1");
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}