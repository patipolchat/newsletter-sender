package handler

import (
	"energy-response-assignment/app/service"
	"energy-response-assignment/config"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Subscriber interface {
	Subscribe(c echo.Context) error
	UnSubscribe(c echo.Context) error
}
type subscriberImpl struct {
	Config            *config.Config
	SubscriberService service.Subscriber
}

type SubscribeRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type SubscribeResponse struct {
	Email   string `json:"email"`
	Active  bool   `json:"active"`
	Message string `json:"message"`
}

func (s *subscriberImpl) Subscribe(c echo.Context) error {
	req := new(SubscribeRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("error binding request: %v", err))
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	subscriber, err := s.SubscriberService.Subscribe(c.Request().Context(), req.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resp := &SubscribeResponse{
		Email:   subscriber.Email,
		Active:  subscriber.Active,
		Message: "Subscribe success",
	}

	return c.JSON(http.StatusOK, resp)
}

func (s *subscriberImpl) UnSubscribe(c echo.Context) error {
	req := new(SubscribeRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("error binding request: %v", err))
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	subscriber, err := s.SubscriberService.Unsubscribe(c.Request().Context(), req.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resp := &SubscribeResponse{
		Email:   subscriber.Email,
		Active:  subscriber.Active,
		Message: "UnSubscribe success",
	}

	return c.JSON(http.StatusOK, resp)
}

func NewSubscriber(cfg *config.Config, subscriberService service.Subscriber) Subscriber {
	return &subscriberImpl{
		Config:            cfg,
		SubscriberService: subscriberService,
	}
}
