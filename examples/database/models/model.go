package models

import "github.com/surrealdb/surrealdb.go"

type TestMigration struct {
	surrealdb.Basemodel `table:"test"`

	ID       string `json:"id,default:()"`
	Username string `json:"username,omitempty" sdb:"optional,assert:(len($value) > 0)"`
	Password string `json:"password,omitempty" sdb:"default:()"`
}
