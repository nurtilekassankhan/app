package service

import "app/pkg/repository"

type Service struct {
	Printer
	Textile
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Printer: NewPrinterService(repos.Printer),
		Textile: NewTextileService(repos.Textile),
	}
}
