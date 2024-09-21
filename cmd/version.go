package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const version = "1.0.7"

var banner = `
   _____                         __   ______         __  __ _                 __      
  / ___/__  ____________  ____ _/ /  / ____/___     / / / /(_)___ _________ _/ /____  
  \__ \/ / / / ___/ ___/ / __ '/ /  / / __/ __ \   / /_/ // / __ '/ ___/ __ '/ __/ _ \ 
 ___/ / /_/ / /  / /    / /_/ / /  / /_/ / /_/ /  / __  // / /_/ / /  / /_/ / /_/  __/ 
/____/\__,_/_/  /_/     \__,_/_/   \____/\____/  /_/ /_//_/\__, /_/   \__,_/\__/\___/  
                                                          /____/                       
`

// SurrealDB colors based on their logo's hues
var (
	purple = color.New(color.FgMagenta).SprintFunc()
	blue   = color.New(color.FgCyan).SprintFunc()
	bold   = color.New(color.Bold).SprintFunc()
)

var initCmd = &cobra.Command{
	Use:   "version",
	Short: "Display Surrealgo Migrate version and information",
	Long:  `Print the version of Surrealgo Migrate along with a brief description of its functionality.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Colorize the banner
		fmt.Println(purple(banner))

		// Print the version in blue
		fmt.Printf("Version: %s\n\n", blue(version))

		// Display the description in styled color
		fmt.Println(bold(purple("Surrealgo Migrate")) + " is a powerful tool for managing database migrations with SurrealDB.")
		fmt.Println("It simplifies the process of creating, applying, and rolling back database changes,")
		fmt.Println("ensuring your database schema stays in sync with your application's needs.")

		// Key features
		fmt.Println("\nKey features:")
		fmt.Println(blue("- Create timestamped migration files"))
		fmt.Println(blue("- Apply migrations to update your database"))
		fmt.Println(blue("- Roll back migrations when needed"))
		fmt.Println(blue("- Support for migration surrealql scripts"))
		fmt.Println(blue("- Transactional migrations for data integrity"))
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
