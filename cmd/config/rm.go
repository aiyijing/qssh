package config

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/aiyijing/qssh/pkg/config"
)

var removeCmd = &cobra.Command{
	Use:     "remove [host]",
	Aliases: []string{"rm"},
	Short:   "remove ssh machine",
	Long:    `remove ssh machine`,
	Example: `qssh config rm 192.168.1.1`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		host := args[0]
		if host != "" {
			m, err := config.QsshConfig.Remove(host)
			if err != nil {
				panic(err)
			}
			if m != nil {
				fmt.Printf("remove host %s success!", m.Host)
			}

		}
	},
}
