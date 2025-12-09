package fakegosqldriver

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

var _ driver.Driver = &fakeDriver {} // implements driver.Driver

type fakeDriver struct {}

func init () {

	sql.Register ("fake", &fakeDriver {})

} // end init

func (*fakeDriver) Open (sConParams string) (con driver.Conn, err error) { // satisfies driver.Driver

	fmt.Printf ("fakeDriver.Open sConParams=%v\n", sConParams)
	con = &fakeConnection {}
	return

} // end method Open
