package report

type StockReport struct {
	ItemID    int     `json:"item_id"`
	ItemName  string  `json:"item_name"`
	PriceSell float64 `json:"price_sell"`
	Stock     int     `json:"stock"`
}

type DailySalesReport struct {
	Date       string  `json:"date"`
	TotalSales float64 `json:"total_sales"`
}

type TopSellingItem struct {
	ItemID    int    `json:"item_id"`
	ItemName  string `json:"item_name"`
	TotalSold int    `json:"total_sold"`
}
