package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/aiyijing/qssh/pkg/config"
	"github.com/aiyijing/qssh/pkg/ssh"
	"github.com/aiyijing/qssh/pkg/util"
)

var (
	password string
	port     int
	key      string
)

var sshCmd = &cobra.Command{
	Use:     "ssh",
	Short:   "Connect to an SSH host",
	Long:    `Connect to an SSH host using the provided configurations.`,
	Example: "qssh ssh root@192.168.1.1 -p 22 -k ~/.ssh/id_rsa",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		sshArgs := args[0]
		user, host := util.ParseConnArgs(sshArgs)
		if user == "" {
			user = "root"
		}
		m, err := config.QsshConfig.Get(host)
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
	sshCmd.Flags().IntVarP(&port, "port", "p", 22, "ssh port")
	sshCmd.Flags().StringVarP(&password, "password", "P", "", "ssh password")
	sshCmd.Flags().StringVarP(&key, "key", "k", "~/.ssh/id_rsa", "ssh private key")
}
