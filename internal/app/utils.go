package app

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	migration_table = "surreal_migrations:initial"
)

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

func defaultConfig() map[string]interface{} {
	return map[string]interface{}{
		"database": map[string]interface{}{
			"connection": map[string]string{
				"user":      "root",
				"password":  "root",
				"name":      "root",
				"namespace": "root",
			},
		},
		"folders": map[string]interface{}{
			"database": map[string]string{
				"migrations": "database/migrations",
				"events":     "database/events",
			},
		},
		"networks": map[string]interface{}{
			"websocket": map[string]string{
				"endpoint": "ws://localhost:8000/rpc",
			},
		},
	}
}
