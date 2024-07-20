package cmd

import (
	"github.com/spf13/cobra"

	"github.com/aiyijing/qssh/cmd/config"
)

var RootCmd = &cobra.Command{
	Use:   "qssh",
	Short: "qssh quickly connects and manages SSH hosts",
	Long:  `qssh quickly connects and manages SSH hosts`,
}

func init() {
	RootCmd.AddCommand(config.ConfigCmd)
	RootCmd.AddCommand(sshCmd)
	RootCmd.AddCommand(execCommand)
}
