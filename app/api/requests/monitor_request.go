package requests

import "github.com/brunobotter/site-sentinel/api/http"

type CreateMonitorTargetRequest struct {
	http.HttpRequest
	URL            string `json:"url"`
	Name           string `json:"name"`
	ExpectedStatus int    `json:"expectedStatus"`
	TimeoutMs      int64  `json:"timeoutMs"`
	IntervalMs     int64  `json:"intervalMs"`
	Retries        int    `json:"retries"`
	IsActive       bool   `json:"isActive"`
}

type ListMonitorResultsRequest struct {
	http.HttpRequest
}

type RunBatchCheckRequest struct {
	http.HttpRequest
}
