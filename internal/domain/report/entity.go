package report

type StockReport struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	PriceSell float64 `json:"price_sell"`
	Stock     int     `json:"stock"`
}

type DailySalesReport struct {
	Date       string  `json:"date"`
	TotalSales float64 `json:"total_sales"`
}

type TopSellingItem struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	TotalSold int    `json:"total_sold"`
}
