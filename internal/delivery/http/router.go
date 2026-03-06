package http

import (
	"net/http"

	"github.com/seuuser/go-site-monitor/internal/service"
)

func NewRouter(siteService *service.SiteService) http.Handler {
	mux := http.NewServeMux()
	handler := NewSiteHandler(siteService)

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})
	mux.HandleFunc("GET /sites", handler.ListSites)
	mux.HandleFunc("POST /sites", handler.CreateSite)
	mux.HandleFunc("DELETE /sites/{id}", handler.DeleteSite)

	return mux
}
