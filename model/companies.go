package model

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"strconv"
	"time"

	"gopkg.in/yaml.v2"
)

import (
	// À vous...
)

/*
Company models a company entry.
*/
type Company struct {
	SIRET       string          `yaml:"siret"`
	Changed     time.Time       `yaml:"changed"`
	TotalSalary currency.Amount `yaml:"-"`
	Staff       []Employee      `yaml:"staff"`
}

/*
MarshalYAML implements the yaml.Marshaler interface
*/
func (c Company) MarshalYAML() (interface{}, error) {
	ready := struct {
		// À vous...
	}{
		// À vous...
	}
	return ready, nil
}

/*
UnmarshalYAML provides the yaml.Unmarshaler interface, to support decoding salary
as a currency.Amount instead of a float64.
*/
func (c *Company) UnmarshalYAML(u func(interface{}) error) (err error) {
	// Use a local anonymous type to avoid re-entering the method.
	tmp := struct {
		// À vous...
	}{}
	err = u(&tmp)
	if err != nil {
		return err
	}
	// À vous...
	c.TotalSalary, err = unmarshalTotalSalary(u)
	return
}

/*
SQLValues returns the fields on a Employee instance as ready to be passed to a SQL statement.
*/
func (c *Company) SQLValues() []interface{} {
	p, _ := c.TotalSalary.MarshalBinary()
	return []interface{}{
		// À vous...
	}
}

/*
Companies is a collection of Company instances.
*/
type Companies []Company

// Comes from a LEFT OUTER JOIN, so all fields are nullable.
type dbStaff struct {
	// À vous...
}
type dbCompany struct {
	// À vous...
}

/*
StoreCompanyV1 stores a Company instance in the database in V1 schema.
*/
func StoreCompanyV1(ctx context.Context, tx *sql.Tx, company Company) error {
	// À vous...
	return nil
}

/*
GetCompanyWithoutStaff retrieves a single company instance by its id, without its staff.
*/
func GetCompanyWithoutStaff(ctx context.Context, tx *sql.Tx, siret string) (Company, error) {
	// À vous... pensez à tx.QueryRowContext
}

/*
GetCompanies retrieves all Company instances from the database in V1 schema.
*/
func GetCompanies(ctx context.Context, tx *sql.Tx) (Companies, error) {
	// Usually don't loop on SQL queries: use a join in a single query instead.

	// À vous...

	return companiesFromDBCompanies(dbCompanies), nil
}

func staffFromDBStaff(dbst []dbStaff) Staff {
	rows := make(Staff, 0, len(dbst))
	// À vous...
	return rows
}

func companiesFromDBCompanies(dbcs []dbCompany) Companies {
	companies := make(Companies, 0, len(dbcs))

	// À vous...
end:
	return companies
}

/*
Show outputs a text representation of a Companies slice, using YAML as the actual format.
*/
func Show(w io.Writer, companies Companies) {
	// À vous...
	fmt.Fprintf(w, "%s\n", y)
}
