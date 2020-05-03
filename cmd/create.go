package cmd

import (
	"github.com/spf13/cobra"
	"tctl/sheety"
)

var task string

var createCmd = &cobra.Command{
	Use:   "create",
	Run: func(cmd *cobra.Command, args []string) { 
		sheety.CreateTask(task);
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&task, "task", "t", "", "task to create.")
	createCmd.MarkFlagRequired("task")
}