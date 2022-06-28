package main

import (
	"context"
	"fmt"
	"log"
	"os"

	atlas "ariga.io/atlas/sql/schema"
	"entgo.io/bug/ent"
	"entgo.io/bug/ent/migrate"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/lib/pq" // needed for ent
	_ "github.com/mattn/go-sqlite3"
)

// Config holds db config.
type Config struct {
	Username string
	Password string
	Address  string
	Port     string
	Database string
	SSLMode  string
}

// ToString converts the db config into pg string.
func (c Config) ToString() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", c.Address, c.Port, c.Username, c.Database, c.Password, c.SSLMode)
}

func main() {

	// client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	// if err != nil {
	// 	log.Fatalf("failed opening connection to sqlite: %v", err)
	// }
	cfg := Config{
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Address:  os.Getenv("DB_ADDR"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_DB"),
		SSLMode:  os.Getenv("DB_SSL"),
	}
	client, err := ent.Open("postgres", cfg.ToString())
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	tx, _ := client.Tx(context.Background())

	if err := MigrationsNoSchema(tx.Client().Debug()); err != nil {
		if err := tx.Rollback(); err != nil {
			log.Fatalf("failed rollback migration no schema %v", err)
		}
		log.Fatalf("failed migration no schema %v", err)
	}

	if err := MigrationsWithSchema(tx.Client().Debug()); err != nil {
		if err := tx.Rollback(); err != nil {
			log.Fatalf("failed rollback migration in schema %v", err)
		}
		log.Fatalf("failed migration in schema %v", err)
	}

	tx.Commit()
}

func MigrationsNoSchema(client *ent.Client) error {
	return client.Schema.Create(context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		schema.WithAtlas(true),
	)
}

func MigrationsWithSchema(client *ent.Client) error {
	return client.Schema.Create(context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		schema.WithAtlas(true),
		schema.WithDiffHook(func(next schema.Differ) schema.Differ {
			return schema.DiffFunc(func(current, desired *atlas.Schema) ([]atlas.Change, error) {
				// Before calculating changes.
				changes, err := next.Diff(current, desired)
				if err != nil {
					return nil, err
				}
				changes = append(changes, &atlas.AddSchema{
					S: atlas.New("mytestschema").
						AddAttrs(desired.Attrs...).
						AddTables(desired.Tables...).
						SetCharset("utf8"),
					Extra: []atlas.Clause{&atlas.IfNotExists{}},
				})
				return changes, nil
				// return []atlas.Change{
				// 	&atlas.AddSchema{
				// 		S: atlas.New("mytestschema").
				// 			AddAttrs(desired.Attrs...).
				// 			AddTables(desired.Tables...).
				// 			SetCharset("utf8"),
				// 		Extra: []atlas.Clause{&atlas.IfNotExists{}},
				// 	},
				// }, nil
			})
		}),
	)
}
