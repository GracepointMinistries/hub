package models

// Pagination contains information about pagination
type Pagination struct {
	Cursor int `db:"cursor" json:"cursor"`
	Limit  int `db:"limit" json:"limit"`
}
