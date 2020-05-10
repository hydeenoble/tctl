package cmd

import (
	"github.com/spf13/cobra"
	"tctl/sheety"
)


var deleteCmd = &cobra.Command{
	Use:   "delete",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) { 
		sheety.Deletetask(args[0]);
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}