package command

import "time"

type CreateTargetCommand struct {
	URL            string
	Name           string
	ExpectedStatus int
	Timeout        time.Duration
	Interval       time.Duration
	Retries        int
	IsActive       bool
}
