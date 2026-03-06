package domain

import "time"

type CheckPolicy struct {
	Timeout        time.Duration
	ExpectedStatus int
	Retries        int
	RetryDelay     time.Duration
}

func DefaultCheckPolicy() CheckPolicy {
	return CheckPolicy{
		Timeout:        3 * time.Second,
		ExpectedStatus: 200,
		Retries:        1,
		RetryDelay:     500 * time.Millisecond,
	}
}
