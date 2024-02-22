package app

import (
	"github.com/spf13/viper"
	"github.com/surrealdb/surrealdb.go"
)

func (mg *Migrator) getDatabase() error {
	// Initialize Viper and read the configuration file

	// Initialize the SurrealDB connection
	db, err := surrealdb.New(viper.GetString(mg.Endpoint))
	if err != nil {
		return err
	}

	// Sign in to the database
	signInData := map[string]interface{}{
		"user": viper.GetString(mg.DatabaseConfig.User),
		"pass": viper.GetString(mg.DatabaseConfig.Password),
	}
	if _, err := db.Signin(signInData); err != nil {
		return err
	}

	// Use the specified database and namespace
	if _, err := db.Use(viper.GetString(mg.DatabaseConfig.Name), viper.GetString(mg.DatabaseConfig.Namespace)); err != nil {
		return err
	}

	// Assign the database connection to the Migrator struct
	mg.db = db

	return nil
}
