package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)
var createCmd = &cobra.Command{
	Use:   "create",
	Run: func(cmd *cobra.Command, args []string) { 
		fmt.Println("Create Command.")
	},
}