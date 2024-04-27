package echoServer

import (
	"context"
	"energy-response-assignment/util/validator"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Config struct {
	Port         int
	Timeout      time.Duration
	AllowOrigins []string
	BodyLimit    string
	LogLevel     string
}

type Server struct {
	App        *echo.Echo
	EchoConfig *Config
}

func (s *Server) GetLogLevel() log.Lvl {
	switch s.EchoConfig.LogLevel {
	case "DEBUG":
		return log.DEBUG
	case "INFO":
		return log.INFO
	case "WARN":
		return log.WARN
	case "ERROR":
		return log.ERROR
	default:
		fmt.Printf("Invalid log level: %s. Defaulting to DEBUG", s.EchoConfig.LogLevel)
		return log.DEBUG
	}
}

func (s *Server) GetRouter() *echo.Router {
	return s.App.Router()
}

func (s *Server) SetupServer() {
	s.App.Logger.SetLevel(s.GetLogLevel())
	s.App.Validator = validator.GetEchoValidator()
	s.setupMiddleWares()
}

func (s *Server) setupMiddleWares() {
	s.App.Use(middleware.Recover())
	s.App.Use(middleware.Logger())
	s.App.Use(GetTimeOutMiddleware(s.EchoConfig.Timeout))
	s.App.Use(GetCORSMiddleware(s.EchoConfig.AllowOrigins))
	s.App.Use(GetBodyLimitMiddleware(s.EchoConfig.BodyLimit))
}

func (s *Server) HttpListening() {
	url := fmt.Sprintf(":%d", s.EchoConfig.Port)
	s.setupGracefullyShutdown()
	if err := s.App.Start(url); err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.App.Logger.Panicf("Error: %v", err)
	}
}

func (s *Server) setupGracefullyShutdown() {
	ctx := context.Background()
	quitCh := make(chan os.Signal, 1)
	go func(quitCh chan os.Signal) {
		signal.Notify(quitCh, syscall.SIGINT, syscall.SIGTERM)
		<-quitCh
		s.App.Logger.Infof("Shutting down service...")

		if err := s.App.Shutdown(ctx); err != nil {
			s.App.Logger.Fatalf("Error: %s", err.Error())
		}
	}(quitCh)
}

func New(cfg *Config) Server {
	echoApp := echo.New()

	return Server{
		App:        echoApp,
		EchoConfig: cfg,
	}
}
