package tzkt

import (
	"reflect"
)

type Filters struct {
	// Greater than filter mode
	TimestampGt string
	// Greater or equal filter mode
	TimestampGe string
	// Less than filter mode
	TimestampLt string
	// Less or equal filter mode
	TimestampLe string
	// Filters items by operation status (applied, failed, backtracked, skipped).
	TypeEq string
	// Sort results desc
	SortDesc string
	// Sort results ac
	SortAsc string
}

var filterParams = map[string]string{
	"TypeEq":      "status.eq",
	"SortAsc":     "sort",
	"SortDesc":    "sort.desc",
	"TimestampGt": "timestamp.gt",
	"TimestampGe": "timestamp.ge",
	"TimestampLt": "timestamp.lt",
	"TimestampLe": "timestamp.le",
}

func (f *Filters) SetFilter(key, value string) {
	if _, ok := filterParams[key]; ok {
		fValue := reflect.ValueOf(f).Elem()
		fValue.FieldByName(key).SetString(value)
	}
}

type Pagination struct {
	Limit int
	// Elements offset mode Skips specified number of elements.
	OffsetEl int
	// Page offset mode Skips page * limit elements. This is a classic pagination...
	OffsetPg int
	// Cursor offset mode. Skips all elements with the cursor before (including) the specified value.
	OffsetCr int
}

var paginationParams = map[string]string{
	"Limit":    "limit",
	"OffsetEl": "offset.el",
	"OffsetPg": "offset.pg",
	"OffsetCr": "offset.cr",
}

func (p *Pagination) SetPagination(key, value string) {
	if _, ok := paginationParams[key]; ok {
		fValue := reflect.ValueOf(p).Elem()
		fValue.FieldByName(key).SetString(value)
	}
}
