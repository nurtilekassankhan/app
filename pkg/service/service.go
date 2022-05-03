package service

import "app/pkg/repository"

type Service struct {
	Printer
	Textile
	Paper
	Paint
	Shirt
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Printer: NewPrinterService(repos.Printer),
		Textile: NewTextileService(repos.Textile),
		Paper:   NewPaperService(repos.Paper),
		Paint:   NewPaintService(repos.Paint),
		Shirt:   NewShirtService(repos.Shirt),
	}
}
