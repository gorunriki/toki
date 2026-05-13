package item

import "time"

type Item struct {
	ID int `json:"id"`

	Name string `json:"name"`

	SKU string `json:"sku"`

	Barcode *string `json:"barcode"`

	PriceSell float64 `json:"price_sell"`

	PriceBuy float64 `json:"price_buy"`

	CreatedAt time.Time `json:"created_at"`

	UpdatedAt time.Time `json:"updated_at"`
}
