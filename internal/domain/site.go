package domain

type Site struct {
	ID              string `json:"id"`
	URL             string `json:"url"`
	IntervalSeconds int    `json:"interval_seconds"`
}
