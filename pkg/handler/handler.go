package handler

import (
	"github.com/dkshi/test-go/pkg/middleware"
	"github.com/dkshi/test-go/pkg/routes"
	"github.com/dkshi/test-go/pkg/service"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) InitRoutes() *echo.Echo {
	mux := echo.New()

	mux.Use(echo.WrapMiddleware(middleware.CheckAdmin))

	mux.GET("/january2025", func(c echo.Context) error {
		return routes.JanuaryTwentyFive(c, h.service.CountDays)
	})

	return mux
}
