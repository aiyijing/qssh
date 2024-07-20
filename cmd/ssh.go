package cmd

import (
	"fmt"
	usr "os/user"

	"github.com/spf13/cobra"

	"github.com/aiyijing/qssh/pkg/config"
	"github.com/aiyijing/qssh/pkg/ssh"
	"github.com/aiyijing/qssh/pkg/util"
)

var (
	password string
	port     int
	key      string
	index    int
)

var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Connect to a machine",
	Long:  `Connect to a machine using the provided config.`,
	Example: `
# Connect to a machine using the private key
qssh ssh root@192.168.1.1 -p 22 -k ~/.ssh/id_rsa

# Connect to a machine using index of config
qssh ssh -i 0
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if index != -1 && len(args) != 0 {
			return fmt.Errorf("cannot specify both index and host args, got %q", args)
		}
		if index == -1 && len(args) == 0 {
			return fmt.Errorf("must specify target host when index is not used")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var (
			user string
			m    *config.Machine
			err  error
		)
		if len(args) != 0 {
			user, host = util.ParseSSHURL(args[0])
			m, err = config.QSSHConfig.Get(host)
		} else {
			m, err = config.QSSHConfig.GetMachineByIndex(index)
			if err != nil {
				fmt.Printf("%v\n", err)
				return
			}
		}
		if user == "" {
			u, _ := usr.Current()
			user = u.Name
		}
		if err == nil {
			user = m.User
			password = m.Password
			host = m.Host
			port = m.Port
			key = m.Key
		}
		client := ssh.NewClient(user, password, host, port, key)
		err = client.Shell()
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	},
}

func init() {
	sshCmd.Flags().IntVarP(&port, "port", "p", 22, "port")
	sshCmd.Flags().StringVarP(&password, "password", "P", "", "password")
	sshCmd.Flags().StringVarP(&key, "key", "k", "~/.ssh/id_rsa", "private key path")
	sshCmd.Flags().IntVarP(&index, "index", "i", -1, "connect host by index of the machine in config")
}
