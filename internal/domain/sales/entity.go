package sales

type Sales struct {
	ID           int
	CustomerName string      `json:"customer_name"`
	TotalAmount  float64     `json:"total_amount"`
	CreatedBy    int         `json:"created_by"`
	Items        []SalesItem `json:"items"`
}

type SalesItem struct {
	ItemID    int     `json:"item_id"`
	Qty       int     `json:"qty"`
	PriceSell float64 `json:"price_sell"`
}
