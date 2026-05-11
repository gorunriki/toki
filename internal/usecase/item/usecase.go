package item

import (
	"context"

	domain "toki/internal/domain/item"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Usecase interface {
	Create(ctx context.Context, it *domain.Item) error
	FindAll(ctx context.Context) ([]domain.Item, error)
	FindByID(ctx context.Context, id int) (*domain.Item, error)
	Update(ctx context.Context, it *domain.Item) error
	Delete(ctx context.Context, id int) error
}

type itemUsecase struct {
	repo domain.Repository
	db   *pgxpool.Pool
}

func NewUsecase(r domain.Repository, db *pgxpool.Pool) Usecase {
	return &itemUsecase{r, db}
}

func (u *itemUsecase) Create(ctx context.Context, it *domain.Item) error {
	tx, err := u.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := `
	INSERT INTO items (name, sku, barcode, price_sell, price_buy)
	VALUES ($1,$2,$3,$4,$5)
	RETURNING id
	`

	var itemID int
	err = tx.QueryRow(ctx, query,
		it.Name, it.SKU, it.Barcode, it.PriceSell, it.PriceBuy,
	).Scan(&itemID)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx,
		`INSERT INTO stocks (item_id, quantity) VALUES ($1,0)`,
		itemID,
	)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (u *itemUsecase) FindAll(ctx context.Context) ([]domain.Item, error) {
	return u.repo.FindAll(ctx)
}

func (u *itemUsecase) FindByID(ctx context.Context, id int) (*domain.Item, error) {
	return u.repo.FindByID(ctx, id)
}

func (u *itemUsecase) Update(ctx context.Context, it *domain.Item) error {
	return u.repo.Update(ctx, it)
}

func (u *itemUsecase) Delete(ctx context.Context, id int) error {
	return u.repo.Delete(ctx, id)
}
