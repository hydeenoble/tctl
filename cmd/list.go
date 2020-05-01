package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"tctl/sheety"
)
var listCmd = &cobra.Command{
	Use:   "list",
	Run: func(cmd *cobra.Command, args []string) { 
		fmt.Println("List Command.")
		sheety.GetTasks();
	},
}