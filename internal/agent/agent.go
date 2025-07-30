package agent

import (
	"log"
	"time"
)

type Agent struct {
	stopCh chan struct{}
}

func New() *Agent {
	return &Agent{
		stopCh: make(chan struct{}),
	}
}

func (a *Agent) Start() error {
	log.Printf("Starting webstack agent")

	// Simple hello world background process
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	log.Println("Agent running in background - Hello World!")

	for {
		select {
		case <-a.stopCh:
			log.Println("Agent stopped")
			return nil
		case <-ticker.C:
			log.Printf("Hello World! Agent heartbeat at %s", time.Now().Format(time.RFC3339))
		}
	}
}

func (a *Agent) Stop() {
	close(a.stopCh)
}
