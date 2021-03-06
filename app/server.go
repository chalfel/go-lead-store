package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chalfel/go-lead-store/config"
	"github.com/chalfel/go-lead-store/db"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	Server *http.Server
	Router *gin.Engine
	DB     *db.DatabaseConnection
	Logger *zap.Logger
	cfg    *config.ApplicationConfig
}

func NewServer(database *db.DatabaseConnection, cfg *config.ApplicationConfig) *Server {
	if cfg.Runtime.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	s := &Server{
		Router: router,
		DB:     database,
		cfg:    cfg,
	}

	s.SetupLogging()

	s.Server = &http.Server{
		Addr:    cfg.Runtime.HttpPort,
		Handler: s.Router,
	}

	return s
}

func (s *Server) Init() error {
	go func() {
		if err := s.Server.ListenAndServe(); err != nil {
			s.Logger.Info("Server is up and listening")
		}
	}()
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Server.Shutdown(ctx); err != nil {
		s.Logger.Fatal("Server forced to shutdown: ")
		return err
	}

	log.Println("Server exiting")
	return nil
}

func (s *Server) SetupLogging() {
	var logger *zap.Logger

	if s.cfg.Runtime.Env == "production" {
		logger, _ = zap.NewProduction()
	} else {
		logger, _ = zap.NewDevelopment()
	}

	defer logger.Sync()
	s.Logger = logger
}
