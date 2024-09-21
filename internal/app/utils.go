package app

import (
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

type DBConfig struct {
	DB       string
	NS       string
	User     string
	Password string
	SSL      bool
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
