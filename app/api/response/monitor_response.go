package response

import "time"

type MonitorTargetResponse struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	URL            string `json:"url"`
	Method         string `json:"method"`
	ExpectedStatus int    `json:"expectedStatus"`
	TimeoutMs      int64  `json:"timeoutMs"`
	Retries        int    `json:"retries"`
	RetryDelayMs   int64  `json:"retryDelayMs"`
	Active         bool   `json:"active"`
	CreatedAt      string `json:"createdAt"`
	UpdatedAt      string `json:"updatedAt"`
}

type CheckResultResponse struct {
	ID             string `json:"id"`
	TargetID       string `json:"targetId"`
	StatusCode     int    `json:"statusCode"`
	ResponseTimeMs int64  `json:"responseTimeMs"`
	IsUp           bool   `json:"isUp"`
	Error          string `json:"error,omitempty"`
	CheckedAt      string `json:"checkedAt"`
}

func FormatTime(t time.Time) string {
	return t.UTC().Format(time.RFC3339)
}
