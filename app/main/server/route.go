package server

import (
	"github.com/brunobotter/site-sentinel/api/controllers"
	"github.com/brunobotter/site-sentinel/main/server/router"
)

func (s *Server) setupApiRouter(healthController *controllers.HealthHandler, monitorController *controllers.MonitorHandler) {
	var routes router.Router
	s.container.NamedResolve(&routes, "Routes")

	s.echo.GET("/health", healthController.Health)

	routes.Group("/monitor", func(group router.RouteGroup) {
		group.POST("/targets", monitorController.CreateTarget)
		group.GET("/targets", monitorController.ListTargets)
		group.POST("/checks/run", monitorController.RunBatchCheck)
		group.GET("/results", monitorController.ListLatestResults)
	})
}
