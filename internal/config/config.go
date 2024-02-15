package config

import "github.com/surrealdb/surrealdb.go"

type Config struct {
	Db *surrealdb.DB
}

func Initialize() *Config {
	return &Config{}
}
