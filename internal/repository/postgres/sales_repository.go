package postgres

import (
	"context"

	"toki/internal/domain/sales"

	"github.com/jackc/pgx/v5"
)

type salesRepository struct{}

func NewSalesRepository() sales.Repository {
	return &salesRepository{}
}

func (r *salesRepository) CreateSales(
	ctx context.Context,
	tx pgx.Tx,
	s *sales.Sales,
) (int, error) {

	query := `
	INSERT INTO sales (
		customer_name,
		total_amount,
		created_by
	)
	VALUES ($1,$2,$3)
	RETURNING id
	`

	var salesID int

	err := tx.QueryRow(
		ctx,
		query,
		s.CustomerName,
		s.TotalAmount,
		s.CreatedBy,
	).Scan(&salesID)

	return salesID, err
}

func (r *salesRepository) CreateSalesItem(
	ctx context.Context,
	tx pgx.Tx,
	salesID int,
	item sales.SalesItem,
) error {

	query := `
	INSERT INTO sales_items (
		sales_id,
		item_id,
		qty,
		price_sell
	)
	VALUES ($1,$2,$3,$4)
	`

	_, err := tx.Exec(
		ctx,
		query,
		salesID,
		item.ItemID,
		item.Qty,
		item.PriceSell,
	)

	return err
}
