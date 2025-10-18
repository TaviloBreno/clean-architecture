package repository

import (
	"database/sql"
	"time"

	"github.com/devfullcycle/goexpert/desafio-clean-architecture/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price, created_at) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice, order.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) FindAll() ([]*entity.Order, error) {
	rows, err := r.Db.Query("SELECT id, price, tax, final_price, created_at FROM orders ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*entity.Order
	for rows.Next() {
		var order entity.Order
		var createdAt string

		err := rows.Scan(&order.ID, &order.Price, &order.Tax, &order.FinalPrice, &createdAt)
		if err != nil {
			return nil, err
		}

		order.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAt)
		if err != nil {
			order.CreatedAt = time.Now()
		}

		orders = append(orders, &order)
	}

	return orders, nil
}
