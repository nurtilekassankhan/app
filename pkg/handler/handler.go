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
		textiles.POST("/", h.createTextile)
		textiles.GET("/", h.getAllTextiles)
		textiles.POST("/:id", h.getTextileById)
		textiles.PUT("/:id", h.updateTextile)
		textiles.DELETE("/:id", h.deleteTextile)
	}

	papers := router.Group("/papers")
	{
		papers.POST("/", h.createPaper)
		papers.GET("/", h.getAllPapers)
		papers.POST("/:id", h.getPaperById)
		papers.PUT("/:id", h.updatePaper)
		papers.DELETE("/:id", h.deletePaper)
	}

	paints := router.Group("/paints")
	{
		paints.POST("/", h.createPaint)
		paints.GET("/", h.getAllPaints)
		paints.POST("/:id", h.getPaintById)
		paints.PUT("/:id", h.updatePaint)
		paints.DELETE("/:id", h.deletePaint)
	}

	shirts := router.Group("/shirts")
	{
		shirts.POST("/", h.createShirt)
		shirts.GET("/", h.getAllShirts)
		shirts.POST("/:id", h.getShirtById)
		shirts.PUT("/:id", h.updateShirt)
		shirts.DELETE("/:id", h.deleteShirt)
	}

	orders := router.Group("/orders")
	{
		orders.POST("/", h.createOrder)
		orders.GET("/", h.getAllOrders)
		orders.POST("/:id", h.getOrderById)
		orders.PUT("/:id", h.updateOrder)
		orders.DELETE("/:id", h.deleteOrder)
	}
	return router
}
