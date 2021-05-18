package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upInstall, downInstall)
}

func upInstall(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	const createCompanies = `
CREATE TABLE companies (
	C'est à vous
);
`
	const createStaff = `
CREATE TABLE staff (
	C'est à vous
  PRIMARY KEY (C'est à vous),
  FOREIGN KEY (C'est à vous) REFERENCES C'est à vous ON C'est à vous
);`
	err := SimpleRun(tx, []string{createCompanies, createStaff})
	return err
}

func downInstall(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	const dropStaff = "C'est à vous"
	const dropCompanies = "C'est à vous"
	err := SimpleRun(tx, []string{dropStaff, dropCompanies})
	return err
}
