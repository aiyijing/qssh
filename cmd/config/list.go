package config

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/aiyijing/qssh/pkg/config"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "list ssh machines",
	Long:    `list ssh machines`,
	Example: `qssh config list`,
	Run: func(cmd *cobra.Command, args []string) {
		machines, _ := config.QsshConfig.List()
		for i, m := range machines {
			fmt.Printf("%v\t%v\n", i, m.Host)
		}
	},
}
