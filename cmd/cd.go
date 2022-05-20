/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

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
		fmt.Println(waypoint);
	},
}

func init() {
	rootCmd.AddCommand(cdCmd)
	cdCmd.PersistentFlags().StringP("shell", "s", "", "The name of the shell in your configuration file")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
