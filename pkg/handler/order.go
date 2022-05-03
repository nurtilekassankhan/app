package handler

import (
	"app"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type getAllOrdersResponse struct {
	Data []app.Order `json:"data"`
}

func (h *Handler) createOrder(c *gin.Context) {
	var input app.Order
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Order.CreateOrder(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllOrders(c *gin.Context) {
	orders, err := h.services.Order.GetAllOrders()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllOrdersResponse{
		Data: orders,
	})
}

func (h *Handler) getOrderById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	order, err := h.services.Order.GetOrderById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, order)
}

func (h *Handler) updateOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	var input app.Order
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err = h.services.Order.UpdateOrder(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	err = h.services.Order.DeleteOrder(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
