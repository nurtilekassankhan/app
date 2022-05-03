package service

import (
	"app"
	"app/pkg/repository"
)

type Paint interface {
	CreatePaint(paint app.Paint) (int, error)
	GetAllPaints() ([]app.Paint, error)
	GetPaintById(id int) (app.Paint, error)
	UpdatePaint(id int, paint app.Paint) error
	DeletePaint(id int) error
}

type PaintService struct {
	repo repository.Paint
}

func NewPaintService(repo repository.Paint) *PaintService {
	return &PaintService{
		repo: repo,
	}
}

func (s *PaintService) CreatePaint(paint app.Paint) (int, error) {
	return s.repo.CreatePaint(paint)
}

func (s *PaintService) GetAllPaints() ([]app.Paint, error) {
	return s.repo.GetAllPaints()
}

func (s *PaintService) GetPaintById(id int) (app.Paint, error) {
	return s.repo.GetPaintById(id)
}

func (s *PaintService) UpdatePaint(id int, paint app.Paint) error {
	return s.repo.UpdatePaint(id, paint)
}

func (s *PaintService) DeletePaint(id int) error {
	return s.repo.DeletePaint(id)
}
