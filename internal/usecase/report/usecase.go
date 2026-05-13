package report

import (
	"context"

	domain "toki/internal/domain/report"
)

type Usecase interface {
	GetStocks(ctx context.Context) ([]domain.StockReport, error)
	GetDailySales(ctx context.Context) ([]domain.DailySalesReport, error)
	GetTopSelling(ctx context.Context) ([]domain.TopSellingItem, error)
}

type reportUsecase struct {
	repo domain.Repository
}

func NewUsecase(r domain.Repository) Usecase {
	return &reportUsecase{
		repo: r,
	}
}

func (u *reportUsecase) GetStocks(
	ctx context.Context,
) ([]domain.StockReport, error) {
	return u.repo.GetStocks(ctx)
}

func (u *reportUsecase) GetDailySales(
	ctx context.Context,
) ([]domain.DailySalesReport, error) {
	return u.repo.GetDailySales(ctx)
}

func (u *reportUsecase) GetTopSelling(
	ctx context.Context,
) ([]domain.TopSellingItem, error) {
	return u.repo.GetTopSelling(ctx)
}
