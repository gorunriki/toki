package postgres

import (
	"context"

	"toki/internal/domain/stock"

	"github.com/jackc/pgx/v5/pgxpool"
)

type stockRepository struct {
	db *pgxpool.Pool
}

func NewStockRepository(db *pgxpool.Pool) stock.Repository {
	return &stockRepository{db}
}

func (r *stockRepository) GetByItemID(ctx context.Context, itemID int) (int, error) {
	var qty int
	err := r.db.QueryRow(ctx,
		`SELECT quantity FROM stocks WHERE item_id=$1`,
		itemID,
	).Scan(&qty)
	return qty, err
}

func (r *stockRepository) UpdateQuantity(ctx context.Context, itemID int, qty int) error {
	_, err := r.db.Exec(ctx,
		`UPDATE stocks SET quantity=$1 WHERE item_id=$2`,
		qty, itemID,
	)
	return err
}
