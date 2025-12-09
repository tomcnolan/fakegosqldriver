package fakegosqldriver

import (
	"database/sql/driver"
	"reflect"
)

var _ driver.Rows                             = &fakeRows {} // implements driver.Rows
var _ driver.RowsColumnTypeDatabaseTypeName   = &fakeRows {} // implements driver.RowsColumnTypeDatabaseTypeName
var _ driver.RowsColumnTypeLength             = &fakeRows {} // implements driver.RowsColumnTypeLength
var _ driver.RowsColumnTypeNullable           = &fakeRows {} // implements driver.RowsColumnTypeNullable
var _ driver.RowsColumnTypePrecisionScale     = &fakeRows {} // implements driver.RowsColumnTypePrecisionScale
var _ driver.RowsColumnTypeScanType           = &fakeRows {} // implements driver.RowsColumnTypeScanType
var _ driver.RowsNextResultSet                = &fakeRows {} // implements driver.RowsNextResultSet

type fakeRows struct {}

func (rows *fakeRows) Close () (err error) { // satisfies driver.Rows

	return

} // end method Close

func (rows *fakeRows) Next (dest [] driver.Value) (err error) { // satisfies driver.Rows

	return

} // end method Next

func (rows *fakeRows) Columns () (asColumns [] string) { // satisfies driver.Rows

	return

} // end method Columns

func (rows *fakeRows) ColumnTypeDatabaseTypeName (iColumn int) (sOutput string) { // satisfies driver.RowsColumnTypeDatabaseTypeName

	return

} // end method ColumnTypeDatabaseTypeName

func (rows *fakeRows) ColumnTypeLength (iColumn int) (nLength int64, bApplicable bool) { // satisfies driver.RowsColumnTypeLength

	return

} // end method ColumnTypeLength

func (rows *fakeRows) ColumnTypeNullable (iColumn int) (bNullable bool, bApplicable bool) { // satisfies driver.RowsColumnTypeNullable

	return

} // end method ColumnTypeNullable

func (rows *fakeRows) ColumnTypePrecisionScale (iColumn int) (nPrecision int64, nScale int64, bApplicable bool) { // satisfies driver.RowsColumnTypePrecisionScale

	return

} // end method ColumnTypePrecisionScale

func (rows *fakeRows) ColumnTypeScanType (iColumn int) (typ reflect.Type) { // satisfies driver.RowsColumnTypeScanType

	return

} // end method ColumnTypeScanType

func (rows *fakeRows) HasNextResultSet () (bOutput bool) { // satisfies driver.RowsNextResultSet

	return

} // end method HasNextResultSet

func (rows *fakeRows) NextResultSet () (err error) { // satisfies driver.RowsNextResultSet

	return

} // end method NextResultSet
