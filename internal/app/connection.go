package app

import (
	"fmt"

	"github.com/surrealdb/surrealdb.go"
)

const (
	migration_table = "surreal_migrations:initial"
)

func (mg *Migrator) getDatabase() error {
	// Initialize Viper and read the configuration file
	// Initialize the SurrealDB connection
	url := fmt.Sprintf("%s://%s/rpc", mg.secureConnection(), mg.Host)
	db, err := surrealdb.New(url)
	if err != nil {
		return err
	}

	if _, err := db.Signin(map[string]interface{}{
		"user": mg.DatabaseConfig.User,
		"pass": mg.DatabaseConfig.Password,
	}); err != nil {
		return err
	}

	// Use the specified database and namespace
	if _, err := db.Use(mg.DatabaseConfig.NS, mg.DatabaseConfig.DB); err != nil {
		return err
	}

	// Assign the database connection to the Migrator struct
	mg.db = db

	return nil
}

func (mg *Migrator) secureConnection() string {
	if mg.DatabaseConfig.SSL {
		return "wss"
	}

	return "ws"
}
