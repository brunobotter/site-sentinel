package server

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"path"
	"strings"
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

//go:embed public/*
var webAssets embed.FS

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
	s.setupWebRouter()
}

func (s *Server) setupWebRouter() {
	webFS, err := fs.Sub(webAssets, "public")
	if err != nil {
		s.logger.Errorf("failed to load embedded web assets: %v", err)
		return
	}

	h := http.FileServer(http.FS(webFS))
	serveWebAsset := func(c echo.Context, assetPath string) error {
		assetPath = strings.TrimPrefix(assetPath, "/")
		assetPath = path.Clean(assetPath)
		if assetPath == "." || assetPath == "" {
			assetPath = "index.html"
		}

		if strings.HasPrefix(assetPath, "../") {
			assetPath = "index.html"
		}

		if _, openErr := webFS.Open(assetPath); openErr != nil {
			assetPath = "index.html"
		}

		// Avoid net/http FileServer canonical redirect loop when serving index.html
		// from a non-root route (e.g. /app/ -> /index.html -> ./ -> /app/ ...).
		if assetPath == "index.html" {
			c.Request().URL.Path = "/"
		} else {
			c.Request().URL.Path = "/" + assetPath
		}
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	}

	s.echo.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusFound, "/app/")
	})

	s.echo.GET("/app", func(c echo.Context) error {
		return c.Redirect(http.StatusFound, "/app/")
	})

	s.echo.GET("/app/", func(c echo.Context) error {
		return serveWebAsset(c, "index.html")
	})

	s.echo.GET("/app/*", func(c echo.Context) error {
		return serveWebAsset(c, c.Param("*"))
	})
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
