package app

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"github.com/surrealdb/surrealdb.go"
)

type Migrator struct {
	db  *surrealdb.DB `yaml:"-"`
	err error         `yaml:"-"`

	Endpoint       string `yaml:"endpoint"`
	DatabaseConfig struct {
		User      string `yaml:"user"`
		Password  string `yaml:"password"`
		Name      string `yaml:"name"`
		Namespace string `yaml:"namespace"`
	} `yaml:"database"`
	FoldersConfig struct {
		Migrations string `yaml:"migrations"`
		Events     string `yaml:"events"`
	} `yaml:"folders"`
}

// gets yaml configuration file and read it
func Migrate() *Migrator {
	viper.SetConfigName("gosurreal")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	var migrator Migrator
	if err := viper.Unmarshal(&migrator, viper.DecoderConfigOption(func(decoderConfig *mapstructure.DecoderConfig) {
		decoderConfig.TagName = "yaml"
	})); err != nil {
		fmt.Println("Make sure you run 'gosurreal init' ")
	}

	migrator.err = err
	return &migrator
}

// check if table exists
func (mg *Migrator) Initialize() *Migrator {
	err := mg.getDatabase()

	if err != nil {
		log.Fatalf("unable to connect to database reason: %s", err.Error())
	}

	_, err = mg.db.Query(fmt.Sprintf(`
		define table if not exists surreal_migrations;
		
		let $m = select * from surreal_migrations;
		if count(m) = 0 {
		    return create %s content {
		        "updated_at": time::now(),
		        "last_migration_id" : "0",
		        "last_event_id": "0"
		    };
		}
	`, migration_table), nil)

	if err != nil {
		log.Fatalf("unable to start migrations error: %s", err.Error())
	}

	return mg
}

func (mg *Migrator) InitConfig() {
	checkConfig()
}

func (mg *Migrator) New(migration string, folderType string) {
	if mg.err != nil {
		log.Println("unable to get configuration error: ", mg.err.Error())
		return
	}

	folder := mg.FoldersConfig.Migrations
	if folderType == "events" {
		folder = mg.FoldersConfig.Events
	}

	mg.createNewMigration(migration, folder)
}

func (mg *Migrator) Exec(migrationType string, folderType string) {
	if mg.db == nil {
		log.Fatalf("make sure the database is connected")
	}

	folder := mg.FoldersConfig.Migrations
	if folderType == "events" {
		folder = mg.FoldersConfig.Events
	}

	files := []string{}
	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		extenstion := ".up.surql"
		if migrationType != "up" {
			extenstion = ".down.surql"
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), extenstion) {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		log.Fatalf("unable to get files from '%s' reason: %s", folder, err.Error())
	}

	migrations := mg.getMigrations(files)

	migrated := false
	for _, file := range files {
		migration, err := mg.getMigration()
		if err != nil {
			log.Fatalf("unable to get current migrations reason: %s", err.Error())
			return
		}

		current := 0

		if folderType == "events" {
			current, _ = strconv.Atoi(migration.LastEventId)
		} else {
			current, _ = strconv.Atoi(migration.LastMigrationId)
		}

		fileName := filepath.Base(file)
		migrationName := strings.Split(fileName, "_")[0]
		timestamp, _ := strconv.Atoi(migrationName)

		if migrationType == "up" {
			if timestamp > current {
				mg.Migrate(file, migrationName, migrationType, folderType == "events", migrationName)
				migrated = true
			}
		} else if migrationType == "down" {
			if current == 0 {
				log.Println("No available migrations to be migrated down")
				break

			} else {
				if migration.LastMigrationId == migrationName {
					newMigraionName, err := mg.findPreviousMigration(migrations, current)
					if err != nil {
						log.Fatal("unable to migrate down reason: ", err)
						return
					}

					mg.Migrate(file, strconv.Itoa(newMigraionName), migrationType, folderType == "events", migrationName)
					break
				}
			}
		}
	}

	if !migrated && migrationType == "up" {
		log.Println("No available migrations to be migrated up")
	}

}

// query
// extra params is used in down to get the file name where down query is
func (mg *Migrator) Migrate(file, migrationName, migrationType string, events bool, extras ...string) {
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	text := string(content)

	if events {
		if _, err := mg.db.Query(fmt.Sprintf(`
		begin transaction;

		update surreal_migrations:initial merge {
			last_event_id: "%s"
		};

		%s

		commit transaction;
	`, migrationName, text), map[string]string{}); err != nil {
			log.Fatalf("unable to migrate %s reason: %s", extras[0], err.Error())
			return
		}
		if migrationType == "up" {
			log.Printf("%s event migrated %s successfully \n", migrationName, migrationType)
		} else {
			log.Printf("%s event migrated %s successfully \n", extras[0], migrationType)
		}

	} else {
		if _, err := mg.db.Query(fmt.Sprintf(`
		begin transaction;

		update surreal_migrations:initial merge {
			last_migration_id: "%s"
		};

		%s

		commit transaction;
	`, migrationName, text), map[string]string{}); err != nil {
			log.Fatalf("unable to migrate %s reason: %s", extras[0], err.Error())
			return
		}

		if migrationType == "up" {
			log.Printf("%s migrated %s successfully \n", migrationName, migrationType)
		} else {
			log.Printf("%s migrated %s successfully \n", extras[0], migrationType)
		}
	}

}
