package config

import (
	"fmt"
	osUser "os/user"

	"github.com/spf13/cobra"

	"github.com/aiyijing/qssh/pkg/config"
	"github.com/aiyijing/qssh/pkg/util"
)

var (
	machine = config.Machine{}
	force   bool
)

var addCmd = &cobra.Command{
	Use:     "add [host]",
	Short:   "add ssh machine",
	Long:    `add ssh machine`,
	Example: `qssh config add root@192.168.1.1 -u root -p 22 -P aiyijing -k ~/root/.ssh/id_rsa`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		user, host := util.ParseConnArgs(args[0])
		if user != "" {
			machine.User = user
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
	u, _ := osUser.Current()
	addCmd.Flags().StringVarP(&(machine.User), "user", "u", u.Name, "ssh user")
	addCmd.Flags().StringVarP(&(machine.Password), "password", "P", "", "ssh password")
	addCmd.Flags().IntVarP(&(machine.Port), "port", "p", 22, "ssh port")
	addCmd.Flags().StringVarP(&(machine.Key), "key", "k", "", "ssh key path")
	addCmd.Flags().BoolVarP(&force, "force", "f", false, "force add")
}
