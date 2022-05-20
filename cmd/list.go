/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"sort"

	"github.com/rafaeltab/EasyCD/helper"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all waypoints",
	Run: func(cmd *cobra.Command, args []string) {
		var waypoints = helper.GetWaypoints()

		sort.Slice(waypoints, func(i, j int) bool {
			return waypoints[i].Name < waypoints[j].Name
		})

		fmt.Println("Waypoints:")
		for i := 0; i < len(waypoints); i++ {
			fmt.Println("  ", waypoints[i].Name, waypoints[i].Path)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
 
}
