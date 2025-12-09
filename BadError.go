package fakegosqldriver

import (
	"database/sql/driver"
)

var _ error = &badError {} // implements error

type badError struct {

	msg string

} // end type badError

func (err *badError) Error () string { // satisfies error

	return err.msg

} // end method Error

func (err *badError) Is (target error) bool {

	return target == driver.ErrBadConn

} // end method Is
