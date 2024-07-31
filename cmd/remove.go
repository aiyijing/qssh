package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/aiyijing/qssh/pkg/config"
)

type RemoveOptions struct {
}

func NewRemoveCommand() *cobra.Command {
	var _ = &RemoveOptions{}
	removeCmd := &cobra.Command{
		Use:     "remove [host]",
		Aliases: []string{"rm"},
		Short:   "remove machine",
		Long:    `remove machine`,
		Example: `
# Using the full command
qssh remove 192.168.1.1

# Using the shorthand command
qssh rm 192.168.1.1

`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			host := args[0]
			if host != "" {
				m, err := config.QSSHConfig.Remove(host)
				if err != nil {
					panic(err)
				}
				if m != nil {
					fmt.Printf("remove host %s success!", m.Host)
				}

			}
		},
	}
	return removeCmd
}
