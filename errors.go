package sqlabble

import "fmt"

// ErrRecordNotFound record not found error
type ErrRecordNotFound struct {
	table string
}

func (err ErrRecordNotFound) Error() string {
	return fmt.Sprintf("%s's table record not found", err.table)
}

// NewErrRecordNotFound create an ErrRecordNotFound instance
// - table: table name
func NewErrRecordNotFound(table string) error {
	return ErrRecordNotFound{
		table: table,
	}
}

// ErrFoundMultipleRecords some records found error
type ErrFoundMultipleRecords struct {
	table string
}

func (err ErrFoundMultipleRecords) Error() string {
	return fmt.Sprintf("multiple records of %s's table are found", err.table)
}

// NewErrFoundMultipleRecords create an ErrFoundMultipleRecords instance
// - table: table name
func NewErrFoundMultipleRecords(table string) error {
	return ErrFoundMultipleRecords{
		table: table,
	}
}
