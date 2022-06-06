package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) tables(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodGet:
		orders, err := h.services.Order.GetAllOrders()
		if err != nil {
			log.Fatal(err)
		}
		shirts, err := h.services.Shirt.GetAllShirts()
		if err != nil {
			log.Fatal(err)
		}
		c.HTML(http.StatusOK, "tables.html", gin.H{
			"Orders": orders,
			"Shirts": shirts,
		})
	default:
		c.JSON(http.StatusMethodNotAllowed, "")
		return
	}
}
