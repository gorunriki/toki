package inbound

type Inbound struct {
	ID        int
	Note      string
	Items     []InboundItem
	CreatedBy int
}

type InboundItem struct {
	ItemID   int     `json:"item_id"`
	Qty      int     `json:"qty"`
	PriceBuy float64 `json:"price_buy"`
}
