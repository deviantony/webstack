package main

import (
	"log"

	"github.com/deviantony/webstack/internal/server"
	"github.com/deviantony/webstack/pkg/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	srv := server.New(cfg)

	log.Printf("Starting webstack server on %s:%d", cfg.Server.Host, cfg.Server.Port)
	if cfg.Server.TLS.Enabled {
		log.Printf("TLS enabled with cert: %s", cfg.Server.TLS.CertFile)
	}

	if err := srv.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
