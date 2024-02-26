/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/JohnKinyanjui/surrealgo_migrate/internal/app"

	"github.com/spf13/cobra"
)

// eventCmd represents the event command
var eventCmd = &cobra.Command{
	Use:        "event",
	Short:      "Provides sets of commands to run event migrations like  'new' ,'up' and 'down' ",
	Long:       ``,
	ArgAliases: []string{"new", "up", "down"},
	Run: func(cmd *cobra.Command, args []string) {
		var migrator = app.Migrate()

		if args[0] == "new" {
			if len(args) > 1 {
				migrator.New(args[1], "events")
			} else {
				fmt.Println("The command should be 'surrealgo migration new add_users' ")
			}

		} else if args[0] == "up" {
			migrator.Initialize().Exec("up", "events")
		} else if args[0] == "down" {
			migrator.Initialize().Exec("down", "events")
		}
	},
}

func init() {
	rootCmd.AddCommand(eventCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// eventCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// eventCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
