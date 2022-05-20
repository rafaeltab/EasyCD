/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/rafaeltab/EasyCD/helper"
	"github.com/spf13/cobra"
)

// cdCmd represents the cd command
var cdCmd = &cobra.Command{
	Use:   "cd waypoint_name",
	Short: "Start a new shell at a waypoint",
	Long: `Start a new shell at a waypoint.
Please note this does not actually do the same thing as the cd command.
Changing the working directory can only be done by the shell using the cd command.
This tool only aims to emulate the same functionality.

Example:
  ecd cd home
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("cd called", args)
		var waypointName = args[0]
		var waypoint = helper.GetWaypoint(waypointName)

		var shellName = cmd.Flag("shell").Value.String()

		if shellName == "" {
			shellName = "default"
		}

		var config = helper.GetConfiguration()

		var shell helper.Shell

		for i := 0; i < len(config.Shells); i++ {
			if config.Shells[i].Name == shellName {
				shell = config.Shells[i]
			}
		}

		if shell.Name == "" {
			fmt.Println("Shell not found")
			os.Exit(1)
		}

		var shl = getShell()

		var arguments = append(shl.Arguments, shell.Command)

		for i := 0; i < len(shell.Arguments); i++ {
			if shell.Arguments[i] == "$DIR$" {
				arguments = append(arguments, waypoint.Path)
				continue
			}

			arguments = append(arguments, shell.Arguments[i])
		}

		exec.Command(shl.Name, arguments...).Start()
	},
}

type commandShell struct {
	Name      string
	Arguments []string
}

func getShell() commandShell {
	// We have to manually add shells since we also need to know how to run commands in them
	var shells = []commandShell{
		{
			Name:      "cmd",
			Arguments: []string{"/c", "start"},
		},
		{
			Name:      "powershell",
			Arguments: []string{"-Command"},
		},
		{
			Name:      "pwsh",
			Arguments: []string{"-Command"},
		},
		{
			Name:      "bash",
			Arguments: []string{"-c"},
		},
		{
			Name:      "zsh",
			Arguments: []string{"-c"},
		},
	}

	for i := 0; i < len(shells); i++ {
		if _, err := exec.LookPath(shells[i].Name); err == nil {
			return shells[i]
		}
	}

	return commandShell{}
}

func init() {
	rootCmd.AddCommand(cdCmd)
	cdCmd.PersistentFlags().StringP("shell", "s", "", "The name of the shell in your configuration file")
}
