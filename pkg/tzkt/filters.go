package tzkt

type Filters struct {
	// Greater than filter mode
	TimestampGt string `url:"timestamp.gt,omitempty"`
	// Greater or equal filter mode
	TimestampGe string `url:"timestamp.ge,omitempty"`
	// Less than filter mode
	TimestampLt string `url:"timestamp.lt,omitempty"`
	// Less or equal filter mode
	TimestampLe string `url:"timestamp.le,omitempty"`
	// Filters items by operation status (applied, failed, backtracked, skipped).
	StatusEq string `url:"status.eq,omitempty"`
	// Sort results desc
	SortDesc string `url:"sort.desc,omitempty"`
	// Sort results ac
	SortAsc string `url:"sort,omitempty"`
}

func NewFilters(options ...func(*Filters)) *Filters {
	fltrs := &Filters{}
	for _, o := range options {
		o(fltrs)
	}
	return fltrs
}

func WithStatusEq(status string) func(*Filters) {
	return func(s *Filters) {
		s.StatusEq = status
	}
}

func WithSortAsc(sort string) func(*Filters) {
	return func(s *Filters) {
		s.SortAsc = sort
	}
}

func WithSortDesc(sort string) func(*Filters) {
	return func(s *Filters) {
		s.SortDesc = sort
	}
}

func WithTimestampGe(timestamp string) func(*Filters) {
	return func(s *Filters) {
		s.TimestampGe = timestamp
	}
}

func WithTimestampGt(timestamp string) func(*Filters) {
	return func(s *Filters) {
		s.TimestampGt = timestamp
	}
}

func WithTimestampLt(timestamp string) func(*Filters) {
	return func(s *Filters) {
		s.TimestampLt = timestamp
	}
}

func WithTimestampLe(timestamp string) func(*Filters) {
	return func(s *Filters) {
		s.TimestampLe = timestamp
	}
}

type Pagination struct {
	Limit int `url:"limit,omitempty"`
	// Elements offset mode Skips specified number of elements.
	OffsetEl int `url:"offset.el,omitempty"`
	// Page offset mode Skips page * limit elements. This is a classic pagination...
	OffsetPg int `url:"offset.pg,omitempty"`
	// Cursor offset mode. Skips all elements with the cursor before (including) the specified value.
	OffsetCr int `url:"offset.cr,omitempty"`
}

func NewPagination(options ...func(*Pagination)) *Pagination {
	page := &Pagination{}
	for _, o := range options {
		o(page)
	}
	return page
}

func WithLimit(limit int) func(*Pagination) {
	return func(s *Pagination) {
		s.Limit = limit
	}
}

func WithOffsetEl(index int) func(*Pagination) {
	return func(s *Pagination) {
		s.OffsetEl = index
	}
}

func WithOffsetPage(page int) func(*Pagination) {
	return func(s *Pagination) {
		s.OffsetPg = page
	}
}

func WithOffsetCursor(id int) func(*Pagination) {
	return func(s *Pagination) {
		s.OffsetCr = id
	}
}
