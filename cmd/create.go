package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"tctl/sheety"
)
var createCmd = &cobra.Command{
	Use:   "create",
	Run: func(cmd *cobra.Command, args []string) { 
		fmt.Println("Create Command.")
		sheety.CreateTask("Work on helm");
	},
}