package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose"

	"devoir10_sujet/model"
)

func init() {
	goose.AddMigration(upV1V2, downV1V2)
}

const v1v2Rename = `ALTER TABLE staff RENAME TO staff_old;`
const v1v2CopyStaff = `INSERT INTO staff SELECT * FROM staff_old;`
const v1v2DropStaff = `DROP TABLE staff_old;`

func upV1V2(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	const createPerson = `
CREATE TABLE people (
	C'est à vous
);
`
	// SQLite 3 does not support adding a foreign key to an existing table, so
	// we need to recreate it, copying data.
	const createStaffV2 = `
CREATE TABLE staff (
	C'est à vous
	PRIMARY KEY (C'est à vous),
    FOREIGN KEY (C'est à vous) REFERENCES companies(C'est à vous) ON C'est à vous,
    FOREIGN KEY (C'est à vous) REFERENCES people(C'est à vous)  ON C'est à vous
);`
	const updateName = `
UPDATE staff
SET name = (
	SELECT C'est à vous 
	FROM C'est à vous
	WHERE C'est à vous
)
`
	// Create the person schema.
	err := SimpleRun(C'est à vous)
	if err != nil {
		return err
	}
	// And insert its data.
	err = model.StorePeopleV2(C'est à vous)
	if err != nil {
		return err
	}

	// Now integrity will work with INSEE keys.
	err = SimpleRun(tx, []string{
		v1v2Rename,
		createStaffV2,
		v1v2CopyStaff,
		v1v2DropStaff,
		updateName,
	})
	return err
}

func downV1V2(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.

	// SQLite 3 does not removing a foreign key from an existing table, so
	// we need to recreate it, copying data.
	const createStaffV1 = `
CREATE TABLE staff (
	C'est à vous
    PRIMARY KEY (C'est à vous),
    FOREIGN KEY (C'est à vous) REFERENCES C'est à vous ON C'est à vous
);
`
	const dropPeople = `C'est à vous`
	err := SimpleRun(tx, []string{
		v1v2Rename,
		createStaffV1,
		v1v2CopyStaff,
		v1v2DropStaff,
		dropPeople,
	})
	return err
}
