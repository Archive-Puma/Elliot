package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cosasdepuma/elliot/app/config"
)

// NewServer TODO Doc
func NewServer() *http.Server {
	return &http.Server{
		Addr:           fmt.Sprintf("%s:%d", config.Host, config.Port),
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
