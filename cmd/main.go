package main

import (
	"context"
	"fmt"
	"log"

	atlas "ariga.io/atlas/sql/schema"
	"entgo.io/bug/ent"
	"entgo.io/bug/ent/migrate"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	if err := MigrationsNoSchema(client); err != nil {
		log.Fatalf("failed migration no schema %v", err)
	}

	if err := MigrationsWithSchema(client); err != nil {
		log.Fatalf("failed migration with schema %v", err)
	}
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
				for i, v := range desired.Tables {
					fmt.Println("tbl", i, desired.Tables[i].Name, v.Name)
				}
				changes = append([]atlas.Change{
					&atlas.AddSchema{
						S: atlas.New("mytestschema").
							AddAttrs(desired.Attrs...).
							AddTables(desired.Tables...).
							SetCharset("utf8"),
						Extra: []atlas.Clause{&atlas.IfNotExists{}},
					},
				}, changes...)
				return changes, nil
			})
		}),
	)
}
