package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type TitleStatistics struct {
	PaintVolume   float64
	TextileVolume float64
	PaperVolume   float64
	PrinterCount  int
}

func (h *Handler) dashboard(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodGet:
		t, err := h.getTitleStatistics()
		if err != nil {
			log.Fatal(err)
		}
		c.HTML(http.StatusOK, "dashboard.html", gin.H{
			"TitleStatistics": t,
		})
	default:
		c.JSON(http.StatusMethodNotAllowed, "")
	}
}

func (h *Handler) getTitleStatistics() (TitleStatistics, error) {

	var resp TitleStatistics

	paints, err := h.services.Paint.GetAllPaints()
	if err != nil {
		return resp, err
	}
	for _, paint := range paints {
		resp.PaintVolume += paint.Volume
	}

	textiles, err := h.services.Textile.GetAllTextiles()
	if err != nil {
		return resp, err
	}
	for _, textile := range textiles {
		resp.TextileVolume += textile.Volume
	}

	papers, err := h.services.Paper.GetAllPaper()
	if err != nil {
		return resp, err
	}
	for _, paper := range papers {
		resp.PaperVolume += paper.Volume
	}

	printers, err := h.services.Printer.GetAllPrinter()
	if err != nil {
		return resp, err
	}
	resp.PrinterCount = len(printers)
	return resp, nil
}
