package domain

import (
	"errors"
	"time"

	"github.com/charly3pins/eShop/domain/base"
	"github.com/gofrs/uuid"
)

func (o *Order) AddProduct(id base.Identity, name, currency string, unitPrice, quantity int) error {
	if o.Processed {
		return ErrOrderAlreadyProcessed
	}

	p := newProduct(id, o.ID.UUID, name, currency, unitPrice, quantity)

	if err := p.Validate(); err != nil {
		return err
	}

	o.TotalPrice = o.TotalPrice + (p.UnitPrice * p.Quantity)
	o.Products = append(o.Products, p)

	return nil
}
func newProduct(id base.Identity, orderID uuid.UUID, name, currency string, unitPrice, quantity int) Product {
	return Product{
		Identity:  id,
		OrderID:   orderID,
		Name:      name,
		UnitPrice: unitPrice,
		Currency:  currency,
		Quantity:  quantity,
	}
}

type Product struct {
	base.Identity
	OrderID   uuid.UUID
	Name      string
	UnitPrice int    // TODO refactor for Money Pattern
	Currency  string // TODO refactor for Money Pattern
	Quantity  int
	CreatedAt time.Time
}

func (p Product) Validate() error {
	if p.OrderID.String() == "" {
		return errors.New("order ID cannot be empty")
	}
	if p.Name == "" {
		return errors.New("name cannot be empty")
	}
	if p.UnitPrice == 0 {
		return errors.New("unit price cannot be 0")
	}
	if p.Currency == "" {
		return errors.New("currency cannot be empty")
	}
	if p.Quantity == 0 {
		return errors.New("quantity cannot be 0")
	}

	return nil
}

type ProductRepository interface {
	base.Repository
	Create(o Order) (Order, error)
	Update(o Order) (Order, error)
	GetOrderByID(id base.Identity) (*Order, error)
}
