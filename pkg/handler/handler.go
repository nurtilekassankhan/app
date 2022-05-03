package handler

import (
	"app/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	printers := router.Group("/printers")
	{
		printers.POST("/", h.createPrinter)
		printers.GET("/", h.getAllPrinters)
		printers.GET("/:id", h.getPrinterById)
		printers.PUT("/:id", h.updatePrinter)
		printers.DELETE("/:id", h.deletePrinter)
	}

	textiles := router.Group("/textiles")
	{
		textiles.POST("/")
	}
	return router
}
