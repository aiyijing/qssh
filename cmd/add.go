package cmd

import (
	"fmt"
	"os/user"

	"github.com/spf13/cobra"

	"github.com/aiyijing/qssh/pkg/config"
	"github.com/aiyijing/qssh/pkg/util"
)

type AddOptions struct {
	Machine  config.Machine
	Override bool
}

func NewAddCommand() *cobra.Command {
	var o = &AddOptions{}
	cmd := &cobra.Command{
		Use:   "add [host]",
		Short: "add machine",
		Long:  `add machine`,
		Example: `# Adding a machine with a password
qssh add root@192.168.1.1 -P admin

# Adding a machine with a private key
qssh add 192.168.1.1 -u root -p 322 -k ~/root/.ssh/id_rsa`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			u, host := util.ParseSSHURL(args[0])
			if u != "" {
				o.Machine.User = u
			}
			if host != "" {
				o.Machine.Host = host
			}
			m, err := config.QSSHConfig.Add(&o.Machine, o.Override)
			if err != nil {
				fmt.Printf("%v\n", err)
				return
			}
			fmt.Printf("add host %s success!\n", m.Host)
		},
	}

	u, _ := user.Current()
	cmd.Flags().StringVarP(&o.Machine.User, "user", "u", u.Name, "user")
	cmd.Flags().StringVarP(&o.Machine.Password, "password", "P", "", "password")
	cmd.Flags().IntVarP(&o.Machine.Port, "port", "p", 22, "port")
	cmd.Flags().StringVarP(&o.Machine.Key, "key", "k", "", "private key path")
	cmd.Flags().BoolVarP(&o.Override, "override", "f", false, "override if the host already exists")
	return cmd
}
