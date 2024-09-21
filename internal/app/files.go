package app

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/surrealdb/surrealdb.go"
)

type Migration struct {
	surrealdb.Basemodel `table:"surreal_migrations"`

	Updated         time.Time `json:"created_At"`
	LastEventId     string    `json:"last_event_id"`
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

func (mg *Migrator) createNewMigration(name string, folder string) {
	// Get the current timestamp
	timestamp := time.Now().Unix()

	// Create the filenames
	upFilename := fmt.Sprintf("%s/%d_%s.up.surql", folder, timestamp, name)
	downFilename := fmt.Sprintf("%s/%d_%s.down.surql", folder, timestamp, name)
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

	if _, err := os.Stat(dirName); err != nil {
		if os.IsNotExist(err) {
			log.Printf("Directory does not exist, creating it")
			if err := os.MkdirAll(dirName, os.ModePerm); err != nil {
				log.Printf("Error creating directory: %v", err)
				return err
			}
		} else {
			log.Printf("Error checking directory: %v", err)
			return err
		}
	}

	file, err := os.Create(fileName)
	if err != nil {
		log.Printf("Error creating file: %v", err)
		return err
	}
	defer file.Close()

	return nil
}
