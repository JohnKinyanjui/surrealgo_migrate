package app

import (
	"github.com/surrealdb/surrealdb.go"
)

func (mg *Migrator) getDatabase() error {
	// Initialize Viper and read the configuration file
	// Initialize the SurrealDB connection
	db, err := surrealdb.New(mg.Endpoint)
	if err != nil {
		return err
	}

	// Sign in to the database
	signInData := map[string]interface{}{
		"user": mg.DatabaseConfig.User,
		"pass": mg.DatabaseConfig.Password,
	}
	if _, err := db.Signin(signInData); err != nil {
		return err
	}

	// Use the specified database and namespace
	if _, err := db.Use(mg.DatabaseConfig.Namespace, mg.DatabaseConfig.Name); err != nil {
		return err
	}

	// Assign the database connection to the Migrator struct
	mg.db = db

	return nil
}
