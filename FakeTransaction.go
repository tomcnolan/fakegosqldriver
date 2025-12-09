package fakegosqldriver

import (
	"database/sql/driver"
)

var _ driver.Tx     = &fakeTransaction {} // implements driver.Tx

type fakeTransaction struct {}

func (tx *fakeTransaction) Commit () (err error) { // satisfies driver.Tx

	return

} // end method Commit

func (tx *fakeTransaction) Rollback () (err error) { // satisfies driver.Tx

	return

} // end method Rollback
