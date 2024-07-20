package config

import "github.com/spf13/cobra"

var ConfigCmd = &cobra.Command{
	Use:     "config",
	Aliases: []string{"cfg"},
	Short:   "config manage machine ",
	Long:    `config manage machine `,
}

func init() {
	ConfigCmd.AddCommand(addCmd)
	ConfigCmd.AddCommand(listCmd)
	ConfigCmd.AddCommand(removeCmd)
}
