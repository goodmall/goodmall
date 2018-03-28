package migration

import (
	"database/sql"
	"fmt"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up00002, Down00002)
}

func Up00002(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	fmt.Println("up to Up00002")
	return nil
}

func Down00002(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.

	fmt.Println("down to Up00002")
	return nil
}
