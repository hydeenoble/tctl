package cmd

import (
	"github.com/spf13/cobra"
	"tctl/sheety"
)
var listCmd = &cobra.Command{
	Use:   "list",
	Run: func(cmd *cobra.Command, args []string) { 
		b, _ := cmd.Flags().GetBool("backlog")
		p, _ := cmd.Flags().GetBool("progress")
		d, _ := cmd.Flags().GetBool("done")

		if p {
			sheety.GetTasks("progress");
		}else if b {
			sheety.GetTasks("backlog");
		}else if d {
			sheety.GetTasks("backlog");
		}else{
			sheety.GetTasks("");
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("backlog", "b", false, "Get only pending tasks.")
	listCmd.Flags().BoolP("progress", "p", false, "Get only tasks that are in progress.")
	listCmd.Flags().BoolP("done", "d", false, "Get only task that have been done.")
}