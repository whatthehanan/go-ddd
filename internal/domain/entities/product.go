package entities

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID        uuid.UUID
	Name      string
	Price     float64
	Seller    Seller
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Product) validate() error {
	if p.Name == "" || p.Price <= 0 {
		return errors.New("invalid product details")
	}

	return nil
}

func NewProduct(name string, price float64, seller ValidatedSeller) *Product {
	return &Product{
		ID:        uuid.New(),
		Name:      name,
		Price:     price,
		Seller:    seller.Seller,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
