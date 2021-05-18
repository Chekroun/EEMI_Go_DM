package model

import (
	"context"
	"database/sql"
	"fmt"
)

/*
Person models a Person used in Employee instances within Company instances.
*/
type Person struct {
	// À vous...
}

/*
SQLValues returns the fields on a Person instance as ready to be passed to a SQL statement.
*/
func (p *Person) SQLValues() []interface{} {
	return []interface{}{
		// À vous...
	}
}

/*
People is a collection of Person.
*/
type People []Person

/*
GetPerson retrieves a single Person instance by its INSEE ID.
*/
func GetPerson(ctx context.Context, tx *sql.Tx, insee string) (Person, error) {
	person := Person{}
	// À vous... pensez à tx.QueryRowContext
}

/*
StorePeopleV2 stores a People instance in the v2 database.
*/
func StorePeopleV2(ctx context.Context, tx *sql.Tx, people People) error {
	// À vous...
}
