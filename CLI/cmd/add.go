package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add_comm",
	Short: "Adds a task to the task list in CLI.",
	Run: func(cmd *cobra.Command, args []string) {
		comm := strings.Join(args, " ")
		fmt.Printf("Added \"%s\" to your task list.\n", comm)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
