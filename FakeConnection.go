package fakegosqldriver

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

var _ driver.Conn                 = &fakeConnection {} // implements driver.Conn
var _ driver.ConnBeginTx          = &fakeConnection {} // implements driver.ConnBeginTx
var _ driver.ConnPrepareContext   = &fakeConnection {} // implements driver.ConnPrepareContext
var _ driver.Execer               = &fakeConnection {} // implements driver.Execer
var _ driver.ExecerContext        = &fakeConnection {} // implements driver.ExecerContext
var _ driver.NamedValueChecker    = &fakeConnection {} // implements driver.NamedValueChecker
var _ driver.Pinger               = &fakeConnection {} // implements driver.Pinger
var _ driver.Queryer              = &fakeConnection {} // implements driver.Queryer
var _ driver.QueryerContext       = &fakeConnection {} // implements driver.QueryerContext

type fakeConnection struct {}

func (con *fakeConnection) Prepare (sRequestText string) (stmt driver.Stmt, err error) { // satisfies driver.Conn

	stmt = &fakeStatement {}
	return

} // end method Prepare

func (con *fakeConnection) Close () (err error) { // satisfies driver.Conn

	return

} // end method Close

func (con *fakeConnection) Begin () (driver.Tx, error) { // satisfies driver.Conn

	return nil, errors.New ("Deprecated method driver.Conn.Begin")

} // end method Begin

func (con *fakeConnection) BeginTx (ctx context.Context, opts driver.TxOptions) (tx driver.Tx, err error) { // satisfies driver.ConnBeginTx

	tx = &fakeTransaction {}
	return

} // end method BeginTx

func (con *fakeConnection) PrepareContext (ctx context.Context, sRequestText string) (stmt driver.Stmt, err error) { // satisfies driver.ConnPrepareContext

	stmt = &fakeStatement {}
	return

} // end method PrepareContext

func (con *fakeConnection) Exec (sRequestText string, aBindValues [] driver.Value) (driver.Result, error) { // satisfies driver.Execer

	return nil, errors.New ("Deprecated method driver.Execer.Exec")

} // end method Exec

func (con *fakeConnection) ExecContext (ctx context.Context, sRequestText string, aBindValues [] driver.NamedValue) (res driver.Result, err error) { // satisfies driver.ExecerContext

	res = &fakeResult {}
	return

} // end method ExecContext

func (con *fakeConnection) CheckNamedValue (pBindValue *driver.NamedValue) (err error) { // satisfies driver.NamedValueChecker

	return

} // end method CheckNamedValue

func (con *fakeConnection) Ping (ctx context.Context) (err error) { // satisfies driver.Pinger

	return

} // end method Ping

func (con *fakeConnection) Query (sRequestText string, aBindValues [] driver.Value) (driver.Rows, error) { // satisfies driver.Queryer

	return nil, errors.New ("Deprecated method driver.Queryer.Query")

} // end method Query

func (con *fakeConnection) QueryContext (ctx context.Context, sRequestText string, aBindValues [] driver.NamedValue) (rows driver.Rows, err error) { // satisfies driver.QueryerContext

	fmt.Printf ("fakeConnection.QueryContext sRequestText=%v\n", sRequestText)

	if strings.HasPrefix (sRequestText, "bad") {
		err = &badError { msg: "fakeConnection.QueryContext simulated error for query: " + sRequestText }
		return
	}

	rows = &fakeRows {}
	return

} // end method QueryContext
