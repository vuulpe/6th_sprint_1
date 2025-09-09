package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	logger *log.Logger
	server *http.Server
}

func New(logger *log.Logger) *Server {
	router := createRouter()

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		logger: logger,
		server: httpServer,
	}
}

func createRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/", handlers.RootHandler)
	router.HandleFunc("/upload", handlers.UploadHandler)

	return router
}

func (s *Server) Start() error {
	s.logger.Printf("Сервер запущен на %s", s.server.Addr)
	return s.server.ListenAndServe()
}
