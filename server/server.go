package server

import (
	"energy-response-assignment/config"
	"energy-response-assignment/util/echoServer"
	"github.com/labstack/echo/v4"
)

type Server interface {
	SetupServer()
	HttpListening()
	GetEchoApp() *echo.Echo
}

type serverImpl struct {
	echoServer.Server
	AppConfig *config.Config
}

func (s *serverImpl) GetEchoApp() *echo.Echo {
	return s.Server.App
}

func NewServer(cfg *config.Config) Server {
	echoConf := &echoServer.Config{
		Port:         cfg.Server.Port,
		Timeout:      cfg.Server.Timeout,
		AllowOrigins: cfg.Server.AllowOrigins,
		BodyLimit:    cfg.Server.BodyLimit,
		LogLevel:     cfg.Server.LogLevel,
	}
	return &serverImpl{
		Server:    echoServer.New(echoConf),
		AppConfig: cfg,
	}
}
