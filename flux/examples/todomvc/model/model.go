package model

// ke: {"package": {"notest": true}}

type Item struct {
	ID       string
	Title    string
	Complete bool
}

type FilterState string

const (
	All       FilterState = "all"
	Active    FilterState = "active"
	Completed FilterState = "completed"
)
