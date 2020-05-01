package cmd

import (
	"github.com/spf13/cobra"
	"tctl/sheety"
)
var listCmd = &cobra.Command{
	Use:   "list",
	Run: func(cmd *cobra.Command, args []string) { 
		b , _ := cmd.Flags().GetBool("backlog")
		p, _ := cmd.Flags().GetBool("progress")
		
		if p {
			sheety.GetTasks("progress");
		}

		if b {
			sheety.GetTasks("backlog");
		}

		
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("backlog", "b", false, "Get only pending tasks")
	listCmd.Flags().BoolP("progress", "p", false, "Get only tasks that are in progress")

}