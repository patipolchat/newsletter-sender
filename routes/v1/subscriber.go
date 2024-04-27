package v1

import (
	"energy-response-assignment/app/handler"
	"energy-response-assignment/app/repository"
	"energy-response-assignment/app/service"
	"energy-response-assignment/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

func SetupSubscriberRoutes(cfg *config.Config, echoApp *echo.Echo, pool *pgxpool.Pool) {
	router := echoApp.Group("/api/v1/subscribers")

	subscriberRepo := repository.NewSubscriber(cfg, pool)
	subscriberService := service.NewSubscriber(cfg, subscriberRepo)
	subscriberHandler := handler.NewSubscriber(cfg, subscriberService)
	router.POST("/subscribe", subscriberHandler.Subscribe)
	router.POST("/unSubscribe", subscriberHandler.UnSubscribe)
}
