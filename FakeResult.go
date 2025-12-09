package fakegosqldriver

import (
	"database/sql/driver"
)

var _ driver.Result = &fakeResult {} // implements driver.Result

type fakeResult struct {}

func (res *fakeResult) LastInsertId () (nId int64, err error) { // satisfies driver.Result

	return

} // end method LastInsertId

func (res *fakeResult) RowsAffected () (nCount int64, err error) { // satisfies driver.Result

	return

} // end method RowsAffected
