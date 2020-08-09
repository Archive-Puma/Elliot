package server

// === IMPORTS ===

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cosasdepuma/elliot/pkg/config"
)

// === PUBLIC METHODS ===

// NewServer generates a new structure associated with the server
func NewServer() *http.Server {
	return &http.Server{
		Addr:           fmt.Sprintf("%s:%d", config.Host, config.Port),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
