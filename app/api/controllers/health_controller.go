package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthHandler struct{}

// NewHealthHandler cria o controller mais simples da API.
//
// Para um júnior: esse handler existe para responder rapidamente se a aplicação está viva.
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// Health responde 200/ok para checagens de liveness.
func (h *HealthHandler) Health(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
