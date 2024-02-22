package app

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/surrealdb/surrealdb.go"
)

type Migration struct {
	surrealdb.Basemodel `table:"surreal_migrations"`

	Updated         time.Time `json:"created_At"`
	LastMigrationId string    `json:"last_migration_id"`
}

func (mg *Migrator) getMigration() (*Migration, error) {
	var migrations Migration
	data, err := mg.db.Select(migration_table)
	if err != nil {
		return nil, err
	}

	err = surrealdb.Unmarshal(data, &migrations)
	if err != nil {
		return nil, err
	}

	return &migrations, nil
}

func (mg *Migrator) createNewMigration(name string) {
	// Get the current timestamp
	timestamp := time.Now().Unix()

	// Create the filenames
	upFilename := fmt.Sprintf("%s/%d_%s.up.surql", mg.FoldersConfig.Migrations, timestamp, name)
	downFilename := fmt.Sprintf("%s/%d_%s.down.surql", mg.FoldersConfig.Migrations, timestamp, name)

	// Create the up file
	if err := ensureDir(upFilename); err != nil {
		fmt.Println("Error creating up file:", err)
		return
	}

	// Create the down file
	if err := ensureDir(downFilename); err != nil {
		fmt.Println("Error creating down file:", err)
		return
	}

	fmt.Println("Created up file:", upFilename)
	fmt.Println("Created down file:", downFilename)
}

func ensureDir(fileName string) error {
	dirName := filepath.Dir(fileName)
	if _, serr := os.Stat(dirName); serr != nil {
		merr := os.MkdirAll(dirName, os.ModePerm)
		if merr != nil {
			return merr
		}
	}

	_, err := os.Create(fileName)
	if err != nil {
		return err
	}

	return nil
}
