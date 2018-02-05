package requests

import "time"

type RequestList struct {
	Fields []string
	Sortby []string
	Order []string
	Query map[string]string
	Limit int64
	Offset int64
	Search string
	OrderBy string
	StartDate time.Time
	EndDate	time.Time
}
