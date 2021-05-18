package model

import (
	"context"
	"database/sql"
	"fmt"

	// À vous...
)

/*
Employee models a Company employee.
*/
type Employee struct {
	// À vous...
}

/*
MarshalYAML implements the yaml.Marshaler interface
*/
func (s Employee) MarshalYAML() (interface{}, error) {
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
func (s *Employee) UnmarshalYAML(u func(interface{}) error) (err error) {
	// Use a local anonymous type to avoid re-entering the method.
	tmp := struct {
		// À vous...
	}{}
	err = u(&tmp)
	if err != nil {
		return err
	}
	// À vous...
	s.Salary, err = unmarshalSalary(u)
	return
}

/*
SQLValues returns the fields on a Employee instance as ready to be passed to a SQL statement.
*/
func (s *Employee) SQLValues() []interface{} {
	p, _ := s.Salary.MarshalBinary()
	return []interface{}{
		// À vous...
	}
}

/*
Staff is a collection of Employee instances.
*/
type Staff []Employee

func storeStaffV1(ctx context.Context, tx *sql.Tx, company Company) error {
	// À vous...
}
