package app

import (
	"energy-response-assignment/config"
	v1 "energy-response-assignment/routes/v1"
	"energy-response-assignment/server"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	Config *config.Config
	Pool   *pgxpool.Pool
	Server server.Server
}

func NewApp(cfg *config.Config, pool *pgxpool.Pool, server server.Server) *App {
	return &App{
		Config: cfg,
		Pool:   pool,
		Server: server,
	}
}

func (a *App) Start() {
	a.Server.SetupServer()
	a.SetupRoute()
	a.Server.HttpListening()
}

func (a *App) SetupRoute() {
	v1.SetupSubscriberRoutes(a.Config, a.Server.GetEchoApp(), a.Pool)
}
