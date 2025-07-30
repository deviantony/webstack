package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/deviantony/webstack/internal/agent"
)

func main() {
	a := agent.New()

	// Handle graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigCh
		log.Println("Received shutdown signal, stopping agent...")
		a.Stop()
		os.Exit(0)
	}()

	log.Printf("Starting webstack agent...")
	if err := a.Start(); err != nil {
		log.Fatalf("Agent failed: %v", err)
	}
}
