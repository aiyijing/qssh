package cmd

import (
	"fmt"
	usr "os/user"

	"github.com/spf13/cobra"

	"github.com/aiyijing/qssh/pkg/config"
	"github.com/aiyijing/qssh/pkg/ssh"
	"github.com/aiyijing/qssh/pkg/util"
)

type SSHOptions struct {
	Password string
	Port     int
	Key      string
	Index    int
}

func NewSSHCommand() *cobra.Command {
	var o = SSHOptions{}
	sshCmd := &cobra.Command{
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
			if o.Index != -1 && len(args) != 0 {
				return fmt.Errorf("cannot specify both index and host args, got %q", args)
			}
			if o.Index == -1 && len(args) == 0 {
				return fmt.Errorf("must specify target host when index is not used")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			var (
				user string
				m    *config.Machine
				host string
				err  error
			)
			if len(args) != 0 {
				user, host = util.ParseSSHURL(args[0])
				m, err = config.QSSHConfig.Get(host)
			} else {
				m, err = config.QSSHConfig.GetMachineByIndex(o.Index)
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
				o.Password = m.Password
				host = m.Host
				o.Port = m.Port
				o.Key = m.Key
			}
			client := ssh.NewClient(user, o.Password, host, o.Port, o.Key)
			err = client.Shell()
			if err != nil {
				fmt.Printf("%v\n", err)
			}
		},
	}
	sshCmd.Flags().IntVarP(&o.Port, "port", "p", 22, "port")
	sshCmd.Flags().StringVarP(&o.Password, "password", "P", "", "password")
	sshCmd.Flags().StringVarP(&o.Key, "key", "k", "~/.ssh/id_rsa", "private key path")
	sshCmd.Flags().IntVarP(&o.Index, "index", "i", -1, "connect host by index of the machine in config")
	return sshCmd
}
