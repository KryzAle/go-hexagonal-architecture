package domain

import (
	"errors"
	"time"

	"github.com/charly3pins/eShop/domain/base"
	"github.com/gofrs/uuid"
)

var (
	ErrOrderAlreadyProcessed = errors.New("order already processed")
)

func NewOrder(id base.Identity, userID uuid.UUID) Order {
	return Order{
		Identity: id,
		UserID:   userID,
	}
}

type Order struct {
	base.Identity
	UserID     uuid.UUID
	TotalPrice int
	Products   []Product
	Processed  bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (o Order) Validate() error {
	if o.UserID.String() == "" {
		return errors.New("user ID cannot be empty")
	}
	if len(o.Products) == 0 {
		return errors.New("order cannot have empty products")
	}

	return nil
}

type OrderRepository interface {
	base.Repository
	Create(o Order) (Order, error)
	Update(o Order) (Order, error)
	GetOrderByID(id base.Identity) (*Order, error)
}
