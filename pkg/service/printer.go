package service

import (
	"app"
	"app/pkg/repository"
)

type Printer interface {
	CreatePrinter(printer app.Printer) (int, error)
	GetAllPrinter() ([]app.Printer, error)
	GetPrinterById(id int) (app.Printer, error)
	UpdatePrinter(id int, printer app.Printer) error
	DeletePrinter(id int) error
}

type PrinterService struct {
	repo repository.Printer
}

func NewPrinterService(repo repository.Printer) *PrinterService {
	return &PrinterService{
		repo: repo,
	}
}

func (s *PrinterService) CreatePrinter(printer app.Printer) (int, error) {
	return s.repo.CreatePrinter(printer)
}

func (s *PrinterService) GetAllPrinter() ([]app.Printer, error) {
	return s.repo.GetAllPrinters()
}

func (s *PrinterService) GetPrinterById(id int) (app.Printer, error) {
	return s.repo.GetPrinterById(id)
}

func (s *PrinterService) UpdatePrinter(id int, printer app.Printer) error {
	return s.repo.UpdatePrinter(id, printer)
}

func (s *PrinterService) DeletePrinter(id int) error {
	return s.repo.DeletePrinter(id)
}
