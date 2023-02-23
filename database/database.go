package database

import (
	"database/sql"
	"fmt"

	"github.com/gobuffalo/packr/v2"
	migrasi "github.com/rubenv/sql-migrate"
)

var (
	DbConnection *sql.DB
)

func DbMigrate(dbParam *sql.DB) {
	migrations := &migrasi.PackrMigrationSource{
		Box: packr.New("migrations", "./sql_migrations"),
	}

	n, errs := migrasi.Exec(dbParam, "postgres", migrations, migrasi.Up)
	if errs != nil {
		panic(errs)
	}

	DbConnection = dbParam
	fmt.Println("Applied ", n, " migrations!")
}
