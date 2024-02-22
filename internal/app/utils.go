package app

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	migration_table = "surreal_migrations:initial"
)

func defaultConfig() map[string]interface{} {
	return map[string]interface{}{
		"endpoint": "ws://localhost:8000/rpc",
		"database": map[string]string{
			"user":      "root",
			"password":  "root",
			"name":      "root",
			"namespace": "root",
		},
		"folders": map[string]string{
			"migrations": "database/migrations",
			"events":     "database/events",
		},
	}
}

func checkConfig() {
	// Check if the file exists
	if _, err := os.Stat("gosurreal.yaml"); os.IsNotExist(err) {
		// File does not exist, create it
		data, err := yaml.Marshal(defaultConfig())
		if err != nil {
			fmt.Println("Error marshaling default config:", err)
			return
		}

		err = os.WriteFile("gosurreal.yaml", data, 0644)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}

		fmt.Println("File 'gosurreal.yaml' created successfully.")
	} else if err != nil {
		fmt.Println("Error checking file:", err)
		return
	} else {
		fmt.Println("File 'gosurreal.yaml' already exists.")
	}
}

func (mg *Migrator) findPreviousMigration(migrations []int, target int) (int, error) {
	for i, num := range migrations {
		if num == target {
			if i == 0 {
				// Target is the first element, no previous exists
				return 0, nil
			}
			return migrations[i-1], nil
		}
	}

	// Target not found in the list
	return 0, fmt.Errorf("migration %d not found in the list", target)
}

func (mg *Migrator) getMigrations(files []string) []int {
	migrations := make([]int, 0)

	for _, file := range files {
		fileName := filepath.Base(file)
		migrationName := strings.Split(fileName, "_")[0]
		timestamp, err := strconv.Atoi(migrationName)
		if err != nil {
			log.Fatalf("currently only supports timestamps")
		}

		migrations = append(migrations, timestamp)
	}

	return migrations
}
