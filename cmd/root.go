/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ecd",
	Short: "A command line tool for creating waypoints for your file system, allowing you to easily go to common directories using a short command.",
	Long: `A command line tool for creating waypoints for your file system, allowing you to easily go to common directories using a short command.

To create a waypoint pointing to the current directory, simply run the command:
  ecd create <waypoint name>

After creating a waypoint, you can go to it using the command:
  ecd cd <waypoint name>

You can also list all waypoints using the command:
  ecd list
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
