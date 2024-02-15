package app

import (
	"github.com/spf13/viper"
	"github.com/surrealdb/surrealdb.go"
)

func (cfg *Migrator) getDatabase() error {
	// Initialize Viper and read the configuration file

	// Initialize the SurrealDB connection
	db, err := surrealdb.New(viper.GetString("networks.websocket.endpoint"))
	if err != nil {
		return err
	}

	// Sign in to the database
	signInData := map[string]interface{}{
		"user": viper.GetString("database.connection.user"),
		"pass": viper.GetString("database.connection.password"),
	}
	if _, err := db.Signin(signInData); err != nil {
		return err
	}

	// Use the specified database and namespace
	if _, err := db.Use(viper.GetString("database.connection.name"), viper.GetString("database.connection.namespace")); err != nil {
		return err
	}

	// Assign the database connection to the Migrator struct
	cfg.db = db

	return nil
}
