package postgres

import (
	"context"

	"toki/internal/domain/item"

	"github.com/jackc/pgx/v5/pgxpool"
)

type itemRepository struct {
	db *pgxpool.Pool
}

func NewItemRepository(db *pgxpool.Pool) item.Repository {
	return &itemRepository{db}
}

func (r *itemRepository) Create(ctx context.Context, it *item.Item) (int, error) {
	query := `
	INSERT INTO items (name, sku, barcode, price_sell, price_buy)
	VALUES ($1,$2,$3,$4,$5)
	RETURNING id
	`
	var id int
	err := r.db.QueryRow(ctx, query,
		it.Name, it.SKU, it.Barcode, it.PriceSell, it.PriceBuy,
	).Scan(&id)
	return id, err
}

func (r *itemRepository) FindAll(ctx context.Context) ([]item.Item, error) {
	rows, err := r.db.Query(ctx, `SELECT id,name,sku,barcode,price_sell,price_buy, created_at, updated_at FROM items`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []item.Item
	for rows.Next() {
		var it item.Item
		err := rows.Scan(&it.ID, &it.Name, &it.SKU, &it.Barcode, &it.PriceSell, &it.PriceBuy, &it.CreatedAt, &it.UpdatedAt)
		if err != nil {
			return nil, err
		}
		items = append(items, it)
	}
	return items, nil
}

func (r *itemRepository) FindByID(ctx context.Context, id int) (*item.Item, error) {
	row := r.db.QueryRow(ctx,
		`SELECT id,name,sku,barcode,price_sell,price_buy,created_at,updated_at FROM items WHERE id=$1`, id)

	var it item.Item
	err := row.Scan(&it.ID, &it.Name, &it.SKU, &it.Barcode, &it.PriceSell, &it.PriceBuy, &it.CreatedAt, &it.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &it, nil
}

func (r *itemRepository) Update(
	ctx context.Context,
	it *item.Item,
) error {

	query := `
	UPDATE items
	SET
		name = $1,
		sku = $2,
		barcode = $3,
		price_sell = $4,
		price_buy = $5,
		updated_at = NOW()
	WHERE id = $6
	`

	_, err := r.db.Exec(
		ctx,
		query,
		it.Name,
		it.SKU,
		it.Barcode,
		it.PriceSell,
		it.PriceBuy,
		it.ID,
	)

	return err
}

func (r *itemRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, `DELETE FROM items WHERE id=$1`, id)
	return err
}
