package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "qssh",
	Short: "qssh quickly connects and manages ssh machines",
	Long:  `qssh quickly connects and manages ssh machines`,
}

func init() {
	RootCmd.AddCommand(NewSSHCommand())
	RootCmd.AddCommand(NewRunCommand())
	RootCmd.AddCommand(NewAddCommand())
	RootCmd.AddCommand(NewListCommand())
	RootCmd.AddCommand(NewRemoveCommand())
	RootCmd.AddCommand(NewCopyCommand())
}
