package http

import (
	"encoding/json"
	"net/http"

	"github.com/seuuser/go-site-monitor/internal/service"
)

type SiteHandler struct {
	siteService *service.SiteService
}

type createSiteRequest struct {
	URL             string `json:"url"`
	IntervalSeconds int    `json:"interval_seconds"`
}

func NewSiteHandler(siteService *service.SiteService) *SiteHandler {
	return &SiteHandler{siteService: siteService}
}

func (h *SiteHandler) ListSites(w http.ResponseWriter, _ *http.Request) {
	sites, err := h.siteService.ListSites()
	if err != nil {
		h.writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.writeJSON(w, http.StatusOK, sites)
}

func (h *SiteHandler) CreateSite(w http.ResponseWriter, r *http.Request) {
	var req createSiteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeError(w, http.StatusBadRequest, "invalid json body")
		return
	}

	site, err := h.siteService.CreateSite(req.URL, req.IntervalSeconds)
	if err != nil {
		h.writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	h.writeJSON(w, http.StatusCreated, site)
}

func (h *SiteHandler) DeleteSite(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := h.siteService.DeleteSite(id); err != nil {
		h.writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *SiteHandler) writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func (h *SiteHandler) writeError(w http.ResponseWriter, status int, message string) {
	h.writeJSON(w, status, map[string]string{"error": message})
}
