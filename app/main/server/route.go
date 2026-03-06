package server

import (
	"github.com/brunobotter/site-sentinel/api/controllers"
	"github.com/brunobotter/site-sentinel/main/server/router"
)

func (s *Server) setupApiRouter(healthController *controllers.HealthHandler) {
	var routs router.Router
	s.container.NamedResolve(&routs, "Routes")

	s.echo.GET("/health", healthController.Health)

}
