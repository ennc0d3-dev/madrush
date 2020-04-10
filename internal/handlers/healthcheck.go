package handlers

import (
	"fmt"
	"net/http"

	log "github.com/ennc0d3/madrush/internal/logger"
)

const (
	healthCheckURI  = "/health"
	livenessURI     = "/liveness"
	healthCheckPort = "9112"
)

type HealthCheckService struct {
}

func NewHealthCheckService() (*HealthCheckService, error) {
	return &HealthCheckService{}, nil
}

func (h *HealthCheckService) Start() error {
	http.HandleFunc(healthCheckURI, healthHandler)
	http.HandleFunc(livenessURI, livenessHandler)
	addr := fmt.Sprintf(":%s", healthCheckPort)
	log.Log("Starting health check service")
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		return err
	}
	return nil
}

func (h *HealthCheckService) Stop() error {
	return nil
}
func (h *HealthCheckService) Restart() error {
	return nil
}
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func livenessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
