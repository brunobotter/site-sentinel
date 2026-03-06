package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/brunobotter/site-sentinel/api/middlewares"
	"github.com/brunobotter/site-sentinel/main/adapters"
	appmiddleware "github.com/brunobotter/site-sentinel/main/server/middleware"
	"github.com/brunobotter/site-sentinel/main/server/router"

	"github.com/brunobotter/site-sentinel/infra/logger"
	"github.com/brunobotter/site-sentinel/main/config"
	"github.com/brunobotter/site-sentinel/main/container"
	"github.com/labstack/echo/v4"
	emw "github.com/labstack/echo/v4/middleware"
)

type Server struct {
	container container.Container
	config    *config.Config
	logger    logger.Logger
	echo      *echo.Echo
}

func NewServer(container container.Container) (*Server, error) {
	server := &Server{
		container: container,
		echo:      echo.New(),
	}
	container.Resolve(&server.config)
	container.Resolve(&server.logger)
	server.setup()
	return server, nil
}

func (s *Server) setup() {
	s.echo.HideBanner = true
	s.echo.Use(emw.CORSWithConfig(emw.CORSConfig{
		AllowOriginFunc: func(origin string) (bool, error) {
			return origin != "", nil
		},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))
	var cfg *config.Config
	s.container.Resolve(&cfg)
	var log logger.Logger
	s.container.Resolve(&log)
	s.echo.Use(appmiddleware.RequestLogger(log))
	adapterRouter := adapters.NewEchoRouterAdapter(s.echo)

	s.container.NamedSingleton("Routes", func() router.Router {
		return adapterRouter.Group("", func(group router.RouteGroup) {
			group.Use(middlewares.CommonMiddlewares(log, cfg)...)
		})
	})

	s.container.Call(s.setupApiRouter)
}

func (s *Server) Run(ctx context.Context) {
	go func() {
		address := fmt.Sprintf(":%d", s.config.Server.Port)
		if err := s.echo.Start(address); err != nil && err != http.ErrServerClosed {
			s.echo.Logger.Fatal(err)
		}
	}()
	s.waitForShutdown(ctx)
}

func (s *Server) waitForShutdown(ctx context.Context) {
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.echo.Shutdown(ctx); err != nil {
		s.echo.Logger.Fatal(err)
	}
}
