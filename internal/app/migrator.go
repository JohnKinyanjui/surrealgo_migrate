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

func (mg *Migrator) Initialize() *Migrator {
	err := mg.getDatabase()

	if err != nil {
		log.Fatalf("unable to connect to database reason: %s", err.Error())
	}

	_, err = mg.db.Query(fmt.Sprintf(`
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

func (mg *Migrator) New(migration string) {
	if mg.err != nil {
		log.Println("unable to get configuration error: ", mg.err.Error())
		return
	}

	mg.createNewMigration(migration)
}

func (mg *Migrator) Exec(migrationType string, folder string) {
	if mg.db == nil {
		log.Fatalf("make sure the database is connected")
	}

	files := []string{}
	err := filepath.Walk(mg.FoldersConfig.Migrations, func(path string, info os.FileInfo, err error) error {
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
		log.Fatal("unable to get files reason: ", err)
	}

	migration, err := mg.getMigration()
	if err != nil {
		return
	}
	current, _ := strconv.Atoi(migration.LastMigrationId)
	migrations := mg.getMigrations(files)

	for _, file := range files {
		fileName := filepath.Base(file)
		migrationName := strings.Split(fileName, "_")[0]
		timestamp, _ := strconv.Atoi(migrationName)

		if migrationType == "up" {
			if timestamp > current {
				mg.Migrate(fileName, migrationName, migrationType)
			}
		} else if migrationType == "down" {
			if current != 0 {
				log.Println("No available migrations to be migrated down")
			} else {
				if migration.LastMigrationId == fileName {
					newMigraionName, err := mg.findPreviousMigration(migrations, current)
					if err != nil {
						log.Fatal("unable to migrate down reason: ", err)
						return
					}

					mg.Migrate(fileName, strconv.Itoa(newMigraionName), migrationType)
				}
			}
		}
	}

}

func (mg *Migrator) Migrate(file, migrationName, migrationType string) {
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	text := string(content)

	_, err = mg.db.Query(fmt.Sprintf(`
				begin transaction;

				update surreal_migrations:initial SET last_migration_id = $last_migration_id;

				%s

				commit transaction;
			`, text), map[string]string{
		"last_migration_id": migrationName,
	})
	if err != nil {
		log.Fatalf("unable to migrate %s reason: %s", migrationName, err.Error())
		return
	}

	log.Printf("%s migrated %s successfully \n", migrationName, migrationType)
}
