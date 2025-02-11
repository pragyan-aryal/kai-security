package actor

import (
	"net/http"
	"strconv"
	"time"

	"kai_security/internal/config"
	"kai_security/internal/logger"

	"github.com/gorilla/mux"
)

type HttpServer struct {
	router *mux.Router
}

func NewHttpServer() *HttpServer {
	return &HttpServer{
		router: mux.NewRouter(),
	}
}

func (h *HttpServer) Start() {
	cfg := config.Get()
	log := logger.Get()

	// TODO
	// Setup routes
	// Setup global middleware (i.e, http request middleware)

	server := &http.Server{
		Addr:         ":" + strconv.Itoa(cfg.HttpPort),
		Handler:      h.router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 50 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Info().Int("port", cfg.HttpPort).Msg("Server starting")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal().Err(err).Msg("Could not start http server")
	}
}
