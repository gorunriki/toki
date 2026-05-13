package report

import "context"

type Repository interface {
	GetStocks(ctx context.Context) ([]StockReport, error)
	GetDailySales(ctx context.Context) ([]DailySalesReport, error)
	GetTopSelling(ctx context.Context) ([]TopSellingItem, error)
}
