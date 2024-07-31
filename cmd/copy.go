package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/aiyijing/qssh/pkg/config"
	"github.com/aiyijing/qssh/pkg/ssh"
	"github.com/aiyijing/qssh/pkg/util"
)

type CopyOptions struct {
	Index  int
	Local  string
	Remote string
}

func NewCopyCommand() *cobra.Command {
	o := &CopyOptions{}
	cmd := &cobra.Command{
		Use:     "copy",
		Aliases: []string{"cp"},
		Short:   "copy local file to remote machine",
		Long:    `copy local file to remote machine`,
		Example: `
# Using the full command
qssh copy /path/to/local/file remote_user@remote_host:/path/to/remote/file

# Using the shorthand command
qssh cp /path/to/local/file remote_host:/path/to/remote/file
`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return fmt.Errorf("requires two arguments when index not set")
			}
			o.Local = args[0]
			o.Remote = args[1]
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			_, host, remotePath := util.ParseSSHURLWithPath(o.Remote)
			fmt.Printf("host %s remotePath %s\n", host, remotePath)
			m, err := config.QSSHConfig.Get(host)
			if err != nil {
				fmt.Printf("%v\n", err)
				os.Exit(1)
			}
			client := ssh.NewClient(m.User, m.Password, m.Host, m.Port, m.Key)
			err = client.Upload(o.Local, remotePath)
			if err != nil {
				fmt.Printf("%v\n", err)
				os.Exit(1)
			}
		},
	}
	cmd.Flags().IntVarP(&o.Index, "index", "i", -1, "index of the machine")
	return cmd
}
