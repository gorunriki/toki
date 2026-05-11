package item

import "context"

type Repository interface {
	Create(ctx context.Context, item *Item) (int, error)
	FindAll(ctx context.Context) ([]Item, error)
	FindByID(ctx context.Context, id int) (*Item, error)
	Update(ctx context.Context, item *Item) error
	Delete(ctx context.Context, id int) error
}
