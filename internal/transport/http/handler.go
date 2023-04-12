package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"v002_onelab/configs"
	"v002_onelab/internal/service"

	_ "v002_onelab/docs"
)

type Handler struct {
	config      *configs.Config
	router      *echo.Echo
	Transaction service.ITransaction
}

func New(service *service.Service, conf *configs.Config) *Handler {
	return &Handler{
		config:      conf,
		router:      echo.New(),
		Transaction: service.Transaction,
	}
}

func (h Handler) InitRouter() *echo.Echo {
	h.router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	api := h.router.Group("/api")
	api.GET("/swagger/*", echoSwagger.WrapHandler)

	api.Use(h.authMiddleware)

	tran := api.Group("/buy-book")
	{
		tran.POST("/", h.BuyBook)
		tran.GET("/", h.GetBuyBooks)
		tran.GET("/:id", h.GetBuyBook)
	}
	return h.router
}
