package config

import (
	"fmt"
	"os/user"

	"github.com/spf13/cobra"

	"github.com/aiyijing/qssh/pkg/config"
	"github.com/aiyijing/qssh/pkg/util"
)

var (
	machine = config.Machine{}
	force   bool
)

var addCmd = &cobra.Command{
	Use:   "add [host]",
	Short: "add machine",
	Long:  `add machine`,
	Example: `# Adding a machine with a password
qssh config add root@192.168.1.1 -P admin

# Adding a machine with a private key
qssh config add 192.168.1.1 -u root -p 322 -k ~/root/.ssh/id_rsa`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		u, host := util.ParseConnArgs(args[0])
		if u != "" {
			machine.User = u
		}
		if host != "" {
			machine.Host = host
		}
		m, err := config.QsshConfig.Add(&machine, force)
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}
		fmt.Printf("add host %s success!\n", m.Host)
	},
}

func init() {
	u, _ := user.Current()

	addCmd.Flags().StringVarP(&(machine.User), "user", "u", u.Name, "ssh user")
	addCmd.Flags().StringVarP(&(machine.Password), "password", "P", "", "ssh password")
	addCmd.Flags().IntVarP(&(machine.Port), "port", "p", 22, "ssh port")
	addCmd.Flags().StringVarP(&(machine.Key), "key", "k", "", "ssh key path")
	addCmd.Flags().BoolVarP(&force, "force", "f", false, "force add")
}
