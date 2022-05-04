package handler

import (
	"app"
	"app/pkg/helper"
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
	//получаем данные о футболке
	shirt, err := h.services.Shirt.GetShirtById(input.ShirtId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	//todo: перенести на получение свободного принтера
	printer, err := h.services.GetPrinterById(4)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	//вычисляем сколько ресурсов нужно для выполнения заказа
	reqResources := helper.CalculateRequiredResources(input, shirt, printer)

	//получаем актуальные данные о ресурсах
	paint, err := h.services.GetPaintById(shirt.PaintId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	paper, err := h.services.GetPaperById(shirt.PaperId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	textile, err := h.services.GetTextileById(shirt.TextileId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	//вычисляем сколько у нас есть ресурсов
	availableResources := helper.GetAvailableResources(paint, paper, textile)

	//смотрим, сможем ли мы исполнить заказ
	if helper.IsOrderCanCompleted(reqResources, availableResources) {
		//создаем ордер на выполнение заказа и возвращаем примерный срок выполнения
		id, err := h.services.Order.CreateOrder(input) //todo: реализовать функцию для снятия объема ресурсов после выполнения ордера
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, map[string]interface{}{
			"id":                 id,
			"required_time_hour": reqResources.SpentTime.Hours(),
		})
	} else {
		newErrorResponse(c, http.StatusInternalServerError, "insufficient resources to execute order") //todo: вынести ошибку в константу
		return
	}
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
