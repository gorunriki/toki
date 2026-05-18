package postgres

import (
	"context"

	"toki/internal/domain/report"

	"github.com/jackc/pgx/v5/pgxpool"
)

type reportRepository struct {
	db *pgxpool.Pool
}

func NewReportRepository(db *pgxpool.Pool) report.Repository {
	return &reportRepository{db}
}

func (r *reportRepository) GetStocks(
	ctx context.Context,
) ([]report.StockReport, error) {

	query := `
	SELECT
		items.id,
		items.name,
		items.price_sell,
		stocks.quantity
	FROM stocks
	JOIN items ON items.id = stocks.item_id
	ORDER BY items.name ASC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var reports []report.StockReport

	for rows.Next() {

		var rep report.StockReport

		err := rows.Scan(
			&rep.ID,
			&rep.Name,
			&rep.PriceSell,
			&rep.Stock,
		)

		if err != nil {
			return nil, err
		}

		reports = append(reports, rep)
	}

	return reports, nil
}

func (r *reportRepository) GetDailySales(
	ctx context.Context,
) ([]report.DailySalesReport, error) {

	query := `
	SELECT
		TO_CHAR(created_at, 'YYYY-MM-DD') as sales_date,
		COALESCE(SUM(total_amount),0)
	FROM sales
	GROUP BY sales_date
	ORDER BY sales_date DESC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var reports []report.DailySalesReport

	for rows.Next() {
		var rep report.DailySalesReport

		err := rows.Scan(
			&rep.Date,
			&rep.TotalSales,
		)

		if err != nil {
			return nil, err
		}

		reports = append(reports, rep)
	}

	return reports, nil
}

func (r *reportRepository) GetTopSelling(
	ctx context.Context,
) ([]report.TopSellingItem, error) {

	query := `
	SELECT
		items.id,
		items.name,
		COALESCE(SUM(sales_items.qty),0) as total_sold
	FROM sales_items
	JOIN items ON items.id = sales_items.item_id
	GROUP BY items.id, items.name
	ORDER BY total_sold DESC
	LIMIT 10
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var reports []report.TopSellingItem

	for rows.Next() {
		var rep report.TopSellingItem

		err := rows.Scan(
			&rep.ID,
			&rep.Name,
			&rep.TotalSold,
		)

		if err != nil {
			return nil, err
		}

		reports = append(reports, rep)
	}

	return reports, nil
}
