package config

import "github.com/spf13/cobra"

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "config mange ssh machine ",
	Long:  `config mange ssh machine `,
}

func init() {
	ConfigCmd.AddCommand(addCmd)
	ConfigCmd.AddCommand(listCmd)
	ConfigCmd.AddCommand(removeCmd)
}
