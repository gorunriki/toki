package sales

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Repository interface {
	CreateSales(ctx context.Context, tx pgx.Tx, sales *Sales) (int, error)
	CreateSalesItem(ctx context.Context, tx pgx.Tx, salesID int, item SalesItem) error
}
