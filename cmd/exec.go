package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/aiyijing/qssh/pkg/config"
	"github.com/aiyijing/qssh/pkg/ssh"
	"github.com/aiyijing/qssh/pkg/util"
)

var (
	ignoreRange string
	host        string
)

var execCommand = &cobra.Command{
	Use:   "exec [script]",
	Short: "Execute commands on SSH hosts",
	Long:  `Execute commands or scripts on SSH hosts using batch processing.`,
	Example: `
	qssh exec "uname -r" --ignore-range 0-1 --host 192.168.1.101
`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var (
			err    error
			rgs    util.Ranges
			script = args[0]
		)
		if ignoreRange != "" {
			rgs, err = util.ParseRanges(ignoreRange)
			if err != nil {
				fmt.Printf("%v\n", err)
				os.Exit(1)
			}
		}
		machines := listMachinesByRange(rgs, host)
		batchExec(script, machines)
	},
}

func init() {
	execCommand.Flags().StringVarP(&ignoreRange, "ignore-range", "i", "", "ignore ssh machine range")
	execCommand.Flags().StringVarP(&host, "host", "H", "", "special ssh machine")
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
	ms, _ := config.QsshConfig.List()
	for i, m := range ms {
		if rgs.Contain(i) && m.Host != specialHost {
			continue
		}
		machines[i] = m
	}
	return machines
}
