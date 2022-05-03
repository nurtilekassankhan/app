package helper

import (
	"app"
	"math"
	"time"
)

func CalculateRequiredResources(order app.Order, shirt app.Shirt, printer app.Printer) app.RequiredResources {
	var result app.RequiredResources
	result.PaintVolume = order.Amount * shirt.PaintVolume
	result.PaperVolume = order.Amount * shirt.PaperVolume
	result.TextileVolume = order.Amount * shirt.TextileVolume
	result.SpentTime = time.Duration(int(math.Round(order.Amount/printer.Efficiency))) * time.Hour
	return result
}

func IsOrderCanCompleted(requiredResources app.RequiredResources, availableResources app.AvailableResources) bool {
	if requiredResources.PaintVolume < availableResources.PaintVolume && requiredResources.PaperVolume < availableResources.PaperVolume && requiredResources.TextileVolume < availableResources.TextileVolume {
		return true
	}
	return false
}

func GetAvailableResources(paint app.Paint, paper app.Paper, textile app.Textile) app.AvailableResources {
	return app.AvailableResources{
		PaintVolume:   paint.Volume,
		PaperVolume:   paper.Volume,
		TextileVolume: textile.Volume,
	}
}
