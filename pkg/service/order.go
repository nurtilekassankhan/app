package service

import (
	"app"
	"app/pkg/repository"
)

type Order interface {
	CreateOrder(order app.Order) (int, error)
	GetAllOrders() ([]app.Order, error)
	GetOrderById(id int) (app.Order, error)
	UpdateOrder(id int, order app.Order) error
	DeleteOrder(id int) error
}

type OrderService struct {
	repo repository.Order
}

func NewOrderService(repo repository.Order) *OrderService {
	return &OrderService{
		repo: repo,
	}
}

func (s *OrderService) CreateOrder(order app.Order) (int, error) {
	return s.repo.CreateOrder(order)
}

func (s *OrderService) GetAllOrders() ([]app.Order, error) {
	return s.repo.GetAllOrders()
}

func (s *OrderService) GetOrderById(id int) (app.Order, error) {
	return s.repo.GetOrderById(id)
}

func (s *OrderService) UpdateOrder(id int, order app.Order) error {
	return s.repo.UpdateOrder(id, order)
}

func (s *OrderService) DeleteOrder(id int) error {
	return s.repo.DeleteOrder(id)
}
