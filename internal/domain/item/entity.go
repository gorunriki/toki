package item

import "time"

type Item struct {
	ID        int
	Name      string
	SKU       string
	Barcode   *string
	PriceSell float64
	PriceBuy  float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
