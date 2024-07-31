package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/aiyijing/qssh/pkg/config"
)

type ListOptions struct {
	output string
}

func NewListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "list machines",
		Long:    `list machines`,
		Example: `
# Using the full command
qssh list

# Using the shorthand command
qssh ls
`,
		Run: func(cmd *cobra.Command, args []string) {
			machines, _ := config.QSSHConfig.List()
			for i, m := range machines {
				fmt.Printf("%v\t%v\n", i, m.Host)
			}
		},
	}
	return cmd
}
