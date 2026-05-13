package stock

import (
	"context"

	domainStock "toki/internal/domain/stock"
)

type Usecase interface {
	GetStock(
		ctx context.Context,
		itemID int,
	) (int, error)
}

type usecase struct {
	repo domainStock.Repository
}

func NewUsecase(
	repo domainStock.Repository,
) Usecase {

	return &usecase{
		repo: repo,
	}
}

func (u *usecase) GetStock(
	ctx context.Context,
	itemID int,
) (int, error) {

	return u.repo.GetByItemID(
		ctx,
		itemID,
	)
}
