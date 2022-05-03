package repository

import (
	"app"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Order interface {
	CreateOrder(order app.Order) (int, error)
	GetOrderById(id int) (app.Order, error)
	GetAllOrders() ([]app.Order, error)
	UpdateOrder(id int, order app.Order) error
	DeleteOrder(id int) error
}

type OrderSQLite struct {
	db *sqlx.DB
}

func NewOrderSQLite(db *sqlx.DB) *OrderSQLite {
	return &OrderSQLite{db: db}
}

func (r *OrderSQLite) CreateOrder(order app.Order) (int, error) {
	var id int
	query := fmt.Sprintf("insert into %s (status, shirt_id, amount) values ($1, $2, $3) returning id", ordersTable)
	row := r.db.QueryRow(query, order.Status, order.ShirtId, order.Amount)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *OrderSQLite) GetOrderById(id int) (app.Order, error) {
	var order app.Order
	query := fmt.Sprintf("select * FROM %s WHERE id = $1", ordersTable)
	if err := r.db.Get(&order, query, id); err != nil {
		return order, err
	}
	return order, nil
}

func (r *OrderSQLite) GetAllOrders() ([]app.Order, error) {
	var orders []app.Order
	query := fmt.Sprintf("select * FROM %s", ordersTable)
	if err := r.db.Select(&orders, query); err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderSQLite) UpdateOrder(id int, order app.Order) error {
	query := fmt.Sprintf("update %s set status = $1, shirt_id = $2, amount = $3 where id = $4", ordersTable)
	_, err := r.db.Exec(query, order.Status, order.ShirtId, order.Amount, id)
	return err
}

func (r *OrderSQLite) DeleteOrder(id int) error {
	query := fmt.Sprintf("delete from %s where id = $1", ordersTable)
	_, err := r.db.Exec(query, id)
	return err
}
