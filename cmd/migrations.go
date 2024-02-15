/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/JohnKinyanjui/surrealgo_migrate/internal/app"
	"github.com/spf13/cobra"
)

// migrationsCmd represents the migrations command
var migrationsCmd = &cobra.Command{
	Use:   "migrations",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	ArgAliases: []string{"new", "up", "down"},
	Run: func(cmd *cobra.Command, args []string) {
		var migrator = app.Migrate()

		if args[0] == "new" {
			println("create files")
			if len(args) > 1 {
				migrator.New(args[1])
			} else {
				fmt.Println("The command should be 'surrealgo migration new add_users' ")
			}

		} else if args[0] == "up" {
			migrator.Initialize().Up()
		}
	},
}

func init() {
	rootCmd.AddCommand(migrationsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrationsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrationsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
