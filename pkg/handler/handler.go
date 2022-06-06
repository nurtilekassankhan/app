package handler

import (
	"app/pkg/service"
	"github.com/gin-gonic/gin"
	"html/template"
	"strings"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.SetFuncMap(template.FuncMap{
		"upper": strings.ToUpper,
	})
	router.Static("/assets", "./ui/assets")
	router.LoadHTMLGlob("./ui/html/*.html")

	app := router.Group("/app")
	{
		app.GET("/dashboard", h.dashboard)
		app.GET("/tables", h.tables)
	}

	printers := router.Group("/printers")
	{
		printers.POST("/", h.createPrinter)      //pass
		printers.GET("/", h.getAllPrinters)      //pass
		printers.GET("/:id", h.getPrinterById)   //pass
		printers.PUT("/:id", h.updatePrinter)    //pass
		printers.DELETE("/:id", h.deletePrinter) //pass
	}

	textiles := router.Group("/textiles")
	{
		textiles.POST("/", h.createTextile)      //pass
		textiles.GET("/", h.getAllTextiles)      //pass
		textiles.GET("/:id", h.getTextileById)   //pass
		textiles.PUT("/:id", h.updateTextile)    //pass
		textiles.DELETE("/:id", h.deleteTextile) //pass
	}

	papers := router.Group("/papers")
	{
		papers.POST("/", h.createPaper)      //pass
		papers.GET("/", h.getAllPapers)      //pass
		papers.GET("/:id", h.getPaperById)   //pass
		papers.PUT("/:id", h.updatePaper)    //pass
		papers.DELETE("/:id", h.deletePaper) //pass
	}

	paints := router.Group("/paints")
	{
		paints.POST("/", h.createPaint)      //pass
		paints.GET("/", h.getAllPaints)      //pass
		paints.GET("/:id", h.getPaintById)   //pass
		paints.PUT("/:id", h.updatePaint)    //pass
		paints.DELETE("/:id", h.deletePaint) //pass
	}

	shirts := router.Group("/shirts")
	{
		shirts.POST("/", h.createShirt)      //pass
		shirts.GET("/", h.getAllShirts)      //pass
		shirts.GET("/:id", h.getShirtById)   //pass
		shirts.PUT("/:id", h.updateShirt)    //pass
		shirts.DELETE("/:id", h.deleteShirt) //pass
	}

	orders := router.Group("/orders")
	{
		orders.POST("/", h.createOrder)      //pass
		orders.GET("/", h.getAllOrders)      //pass
		orders.GET("/:id", h.getOrderById)   //pass
		orders.PUT("/:id", h.updateOrder)    //pass
		orders.DELETE("/:id", h.deleteOrder) //pass
	}
	return router
}
