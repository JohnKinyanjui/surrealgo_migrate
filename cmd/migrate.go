/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/JohnKinyanjui/surrealgo_migrate/internal/app"
	"github.com/spf13/cobra"
)

var cfg app.Migrator

// migrationsCmd represents the migrations command
var migrationsCmd = &cobra.Command{
	Use:        "migrate",
	Short:      "Provides sets of commands to run migrations like  'new' ,'up' and 'down' ",
	Long:       ``,
	ArgAliases: []string{"new", "up", "down"},
	Run: func(cmd *cobra.Command, args []string) {
		var migrator = app.Migrate(&cfg)
		if args[0] == "new" {
			if !cmd.Flags().Changed("path") {
				log.Fatal("Error: --path flags are required")
			}

			if len(args) > 1 {
				migrator.New(args[1])
			} else {
				fmt.Println("The command should be 'surrealgo_migrate migrate --path [MIGRATION_FOLDER] new [FILE_NAME]' ")
			}

			return
		}

		if !cmd.Flags().Changed("host") || !cmd.Flags().Changed("db") || !cmd.Flags().Changed("ns") || !cmd.Flags().Changed("user") || !cmd.Flags().Changed("pass") || !cmd.Flags().Changed("path") {
			log.Fatal("Error: --host, --db, --ns, --user --pass and --path flags are required")
		}

		if args[0] == "up" {
			migrator.Initialize().Exec("up")
		} else if args[0] == "down" {
			migrator.Initialize().Exec("down")
		}

		if len(args) > 0 {
			fmt.Println("Running migration:", args[0])
		} else {
			fmt.Println("No migration command provided.")
		}
	},

	// Define required flags

}

func init() {
	rootCmd.AddCommand(migrationsCmd)

	// Here you will define your flags and configuration settings.
	migrationsCmd.Flags().StringVar(&cfg.Path, "path", "", "migration folder is required")

	migrationsCmd.Flags().StringVar(&cfg.Host, "host", "", "SurrealDB server URL (required)")
	migrationsCmd.Flags().StringVar(&cfg.DatabaseConfig.DB, "db", "", "Database name to connect to (required)")
	migrationsCmd.Flags().StringVar(&cfg.DatabaseConfig.NS, "ns", "", "Namespace within the database (required)")
	migrationsCmd.Flags().StringVar(&cfg.DatabaseConfig.User, "user", "", "Username for database authentication (required)")
	migrationsCmd.Flags().StringVar(&cfg.DatabaseConfig.Password, "pass", "", "Password for database authentication (required)")
	migrationsCmd.Flags().BoolVar(&cfg.DatabaseConfig.SSL, "ssl", false, "Enable SSL/TLS for secure connection (default: false)")

	// User      string `yaml:"user"`
	// Password  string `yaml:"password"`
	// Name      string `yaml:"name"`
	// Namespace string `yaml:"namespace"`
	// Mark the flags as required
	// rootCmd.MarkFlagRequired("host")
	migrationsCmd.MarkFlagRequired("path")
	// Execute the command
	// if err := rootCmd.Execute(); err != nil {
	// 	log.Fatal(err)
	// }
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrationsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrationsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
