package entity

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID         string    `json:"id"`
	Price      float64   `json:"price"`
	Tax        float64   `json:"tax"`
	FinalPrice float64   `json:"final_price"`
	CreatedAt  time.Time `json:"created_at"`
}

func NewOrder(id, price, tax string) (*Order, error) {
	order := &Order{}

	if err := order.SetID(id); err != nil {
		return nil, err
	}

	if err := order.SetPrice(price); err != nil {
		return nil, err
	}

	if err := order.SetTax(tax); err != nil {
		return nil, err
	}

	order.CalculateFinalPrice()
	order.CreatedAt = time.Now()

	if err := order.IsValid(); err != nil {
		return nil, err
	}

	return order, nil
}

func (o *Order) SetID(id string) error {
	if id == "" {
		o.ID = uuid.New().String()
	} else {
		o.ID = id
	}
	return nil
}

func (o *Order) SetPrice(price string) error {
	if price == "" {
		return errors.New("price is required")
	}

	var p float64
	if _, err := fmt.Sscanf(price, "%f", &p); err != nil {
		return errors.New("invalid price")
	}

	if p <= 0 {
		return errors.New("price must be greater than 0")
	}

	o.Price = p
	return nil
}

func (o *Order) SetTax(tax string) error {
	if tax == "" {
		return errors.New("tax is required")
	}

	var t float64
	if _, err := fmt.Sscanf(tax, "%f", &t); err != nil {
		return errors.New("invalid tax")
	}

	if t < 0 {
		return errors.New("tax must be greater than or equal to 0")
	}

	o.Tax = t
	return nil
}

func (o *Order) CalculateFinalPrice() {
	o.FinalPrice = o.Price + o.Tax
}

func (o *Order) IsValid() error {
	if o.ID == "" {
		return errors.New("id is required")
	}
	if o.Price <= 0 {
		return errors.New("price must be greater than 0")
	}
	if o.Tax < 0 {
		return errors.New("tax must be greater than or equal to 0")
	}
	return nil
}
