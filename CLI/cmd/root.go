package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task manager is a Command line Interface tool",
}
