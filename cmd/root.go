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
	RootCmd.AddCommand(sshCmd)
	RootCmd.AddCommand(runCmd)
	RootCmd.AddCommand(addCmd)
	RootCmd.AddCommand(listCmd)
	RootCmd.AddCommand(removeCmd)
}
