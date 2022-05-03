package service

import (
	"app"
	"app/pkg/repository"
)

type Paper interface {
	CreatePaper(paper app.Paper) (int, error)
	GetAllPaper() ([]app.Paper, error)
	GetPaperById(id int) (app.Paper, error)
	UpdatePaper(id int, paper app.Paper) error
	DeletePaper(id int) error
}

type PaperService struct {
	repo repository.Paper
}

func NewPaperService(repo repository.Paper) *PaperService {
	return &PaperService{
		repo: repo,
	}
}

func (s *PaperService) CreatePaper(paper app.Paper) (int, error) {
	return s.repo.CreatePaper(paper)
}

func (s *PaperService) GetAllPaper() ([]app.Paper, error) {
	return s.repo.GetAllPapers()
}

func (s *PaperService) GetPaperById(id int) (app.Paper, error) {
	return s.repo.GetPaperById(id)
}

func (s *PaperService) UpdatePaper(id int, paper app.Paper) error {
	return s.repo.UpdatePaper(id, paper)
}

func (s *PaperService) DeletePaper(id int) error {
	return s.repo.DeletePaper(id)
}
