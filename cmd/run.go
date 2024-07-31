package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/aiyijing/qssh/pkg/config"
	"github.com/aiyijing/qssh/pkg/ssh"
	"github.com/aiyijing/qssh/pkg/util"
)

type RunOptions struct {
	IgnoreRange string
	Host        string
}

func NewRunCommand() *cobra.Command {
	var o = &RunOptions{}
	runCmd := &cobra.Command{
		Use:   "run [script]",
		Short: "Execute commands on remote hosts",
		Long:  `Execute commands or scripts on hosts using batch processing.`,
		Example: `
#  Execute commands on remote hosts
qssh run "uname -r" --ignore-range 0-1 --host 192.168.1.101
`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var (
				err    error
				rgs    util.Ranges
				script = args[0]
			)
			if o.IgnoreRange != "" {
				rgs, err = util.ParseRanges(o.IgnoreRange)
				if err != nil {
					fmt.Printf("%v\n", err)
					os.Exit(1)
				}
			}
			machines := listMachinesByRange(rgs, o.Host)
			batchExec(script, machines)
		},
	}
	runCmd.Flags().StringVarP(&o.IgnoreRange, "ignore-range", "i", "", "ignore machine range")
	runCmd.Flags().StringVarP(&o.Host, "host", "H", "", "special host")
	return runCmd
}

func batchExec(script string, machines map[int]*config.Machine) {
	for i, m := range machines {
		fmt.Printf("[%d] %s\n", i, m.Host)
		client := ssh.NewClient(m.User, m.Password, m.Host, m.Port, m.Key)
		result, err := client.Run(script)
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}
		fmt.Printf("%s\n", result)
	}
}

func listMachinesByRange(rgs util.Ranges, specialHost string) map[int]*config.Machine {
	var machines = make(map[int]*config.Machine)
	ms, _ := config.QSSHConfig.List()
	for i, m := range ms {
		if rgs.Contain(i) && m.Host != specialHost {
			continue
		}
		machines[i] = m
	}
	return machines
}
