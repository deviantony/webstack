package server

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/deviantony/webstack/pkg/config"
	"github.com/gin-gonic/gin"
)

//go:embed static/*
var staticFiles embed.FS

type Server struct {
	config *config.Config
	router *gin.Engine
}

func New(cfg *config.Config) *Server {
	if !cfg.Server.TLS.Enabled {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	server := &Server{
		config: cfg,
		router: router,
	}

	server.setupRoutes()
	return server
}

func (s *Server) setupRoutes() {
	// Health check endpoint
	s.router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"service": "webstack-server",
		})
	})

	// API routes group
	api := s.router.Group("/api")
	{
		api.GET("/status", s.handleStatus)
	}

	// Serve static files (React app)
	s.setupStaticFiles()
}

func (s *Server) setupStaticFiles() {
	// Serve React app from embedded files
	staticFS, err := fs.Sub(staticFiles, "static")
	if err != nil {
		// If static files not found, show error
		s.router.GET("/", func(c *gin.Context) {
			c.String(http.StatusInternalServerError, "Static files not found. Please run 'npm run build' in the web directory.")
		})
		return
	}

	// Serve static assets (CSS, JS, etc.)
	assetsFS, err := fs.Sub(staticFS, "assets")
	if err == nil {
		s.router.StaticFS("/assets", http.FS(assetsFS))
	}

	// Serve index.html for all non-API routes (SPA routing)
	s.router.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path

		// Skip API and WebSocket routes
		if len(path) >= 4 && path[:4] == "/api" {
			c.Status(http.StatusNotFound)
			return
		}
		if len(path) >= 3 && path[:3] == "/ws" {
			c.Status(http.StatusNotFound)
			return
		}

		// For all other routes, serve the React app
		data, err := staticFiles.ReadFile("static/index.html")
		if err != nil {
			c.String(http.StatusNotFound, "Page not found")
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", data)
	})
}

func (s *Server) handleStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"service": "webstack-server",
		"version": "0.1.0",
		"status":  "running",
		"config": gin.H{
			"host": s.config.Server.Host,
			"port": s.config.Server.Port,
			"tls":  s.config.Server.TLS.Enabled,
		},
	})
}

func (s *Server) Start() error {
	addr := fmt.Sprintf("%s:%d", s.config.Server.Host, s.config.Server.Port)

	if s.config.Server.TLS.Enabled {
		return s.router.RunTLS(addr, s.config.Server.TLS.CertFile, s.config.Server.TLS.KeyFile)
	}

	return s.router.Run(addr)
}
