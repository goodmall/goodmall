package migration

import (
	"database/sql"
	"fmt"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up00001, Down00001)
}

func Up00001(tx *sql.Tx) error {
	// This code is executed when the migration is applied.

	fmt.Println("up up !")
	return nil
}

func Down00001(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.

	fmt.Println("down down !")
	return nil
}
