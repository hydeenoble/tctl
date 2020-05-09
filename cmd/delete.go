package cmd

import (
	"github.com/spf13/cobra"
	"tctl/sheety"
)


var deleteCmd = &cobra.Command{
	Use:   "delete",
	Run: func(cmd *cobra.Command, args []string) { 
		sheety.Deletetask();
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}