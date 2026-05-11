package stock

import (
	"context"
	"errors"

	"toki/internal/domain/stock"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Usecase interface {
	GetStock(ctx context.Context, itemID int) (int, error)
	DecreaseStock(ctx context.Context, itemID int, qty int) error
}

type stockUsecase struct {
	repo stock.Repository
	db   *pgxpool.Pool
}

func NewUsecase(r stock.Repository, db *pgxpool.Pool) Usecase {
	return &stockUsecase{r, db}
}

func (u *stockUsecase) GetStock(ctx context.Context, itemID int) (int, error) {
	return u.repo.GetByItemID(ctx, itemID)
}

func (u *stockUsecase) DecreaseStock(ctx context.Context, itemID int, qty int) error {
	current, err := u.repo.GetByItemID(ctx, itemID)
	if err != nil {
		return err
	}

	if current < qty {
		return errors.New("stock not enough")
	}

	return u.repo.UpdateQuantity(ctx, itemID, current-qty)
}
