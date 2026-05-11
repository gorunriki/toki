package stock

import "context"

type Repository interface {
	GetByItemID(ctx context.Context, itemID int) (int, error)
	UpdateQuantity(ctx context.Context, itemID int, qty int) error
}
