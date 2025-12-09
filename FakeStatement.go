package fakegosqldriver

import (
	"context"
	"database/sql/driver"
	"errors"
)

var _ driver.ColumnConverter    = &fakeStatement {} // implements driver.ColumnConverter
var _ driver.Stmt               = &fakeStatement {} // implements driver.Stmt
var _ driver.StmtExecContext    = &fakeStatement {} // implements driver.StmtExecContext
var _ driver.StmtQueryContext   = &fakeStatement {} // implements driver.StmtQueryContext

type fakeStatement struct {}

func (stmt *fakeStatement) ColumnConverter (iParam int) driver.ValueConverter { // satisfies driver.ColumnConverter

	return driver.DefaultParameterConverter

} // end method ColumnConverter

func (stmt *fakeStatement) Close () (err error) { // satisfies driver.Stmt

	return

} // end method Close

func (stmt *fakeStatement) NumInput () int { // satisfies driver.Stmt

	return -1 // suppress bind value count checking by sql package

} // end method NumInput

func (stmt *fakeStatement) Exec (args [] driver.Value) (driver.Result, error) { // satisfies driver.Stmt

	return nil, errors.New ("Deprecated method driver.Stmt.Exec")

} // end method Exec

func (stmt *fakeStatement) Query (args [] driver.Value) (driver.Rows, error) { // satisfies driver.Stmt

	return nil, errors.New ("Deprecated method driver.Stmt.Query")

} // end method Query

func (stmt *fakeStatement) ExecContext (ctx context.Context, aBindValues [] driver.NamedValue) (res driver.Result, err error) { // satisfies driver.StmtExecContext

	res = &fakeResult {}
	return

} // end method ExecContext

func (stmt *fakeStatement) QueryContext (ctx context.Context, aBindValues [] driver.NamedValue) (rows driver.Rows, err error) { // satisfies driver.StmtQueryContext

	rows = &fakeRows {}
	return

} // end method QueryContext
