package cmd

import (
	"github.com/spf13/cobra"
	"tctl/sheety"
)
var listCmd = &cobra.Command{
	Use:   "list",
	Run: func(cmd *cobra.Command, args []string) { 
		p , _ := cmd.Flags().GetBool("pending")
		if p {
			sheety.GetTasks("pending");
		}
		
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("pending", "p", false, "Get only pending tasks")
}