/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/rafaeltab/EasyCD/helper"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create waypoint_name",
	Short: "Create a waypoint",
	Long: `Create a waypoint.
By default this command will create the waypoint in the current directory.
	
For example, you can create a waypoint called "home" that points to your current directory using the following command.
	ecd create home`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var waypointName = args[0]
		var directory = cmd.Flag("directory").Value.String()

		if directory == "" {
			var wd, err = os.Getwd()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			directory = wd
		}

		helper.AddWaypoint(waypointName, directory)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.PersistentFlags().StringP("directory", "d", "", "Specify a directory to create the waypoint in")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
