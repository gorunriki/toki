package postgres

import (
	"context"

	"toki/internal/domain/inbound"

	"github.com/jackc/pgx/v5"
)

type inboundRepository struct{}

func NewInboundRepository() inbound.Repository {
	return &inboundRepository{}
}

func (r *inboundRepository) CreateInbound(
	ctx context.Context,
	tx pgx.Tx,
	in *inbound.Inbound,
) (int, error) {

	var inboundID int

	query := `
	INSERT INTO inbounds (note, created_by)
	VALUES ($1,$2)
	RETURNING id
	`

	err := tx.QueryRow(
		ctx,
		query,
		in.Note,
		in.CreatedBy,
	).Scan(&inboundID)

	return inboundID, err
}

func (r *inboundRepository) CreateInboundItem(
	ctx context.Context,
	tx pgx.Tx,
	inboundID int,
	item inbound.InboundItem,
) error {

	query := `
	INSERT INTO inbound_items (
		inbound_id,
		item_id,
		qty,
		price_buy
	)
	VALUES ($1,$2,$3,$4)
	`

	_, err := tx.Exec(
		ctx,
		query,
		inboundID,
		item.ItemID,
		item.Qty,
		item.PriceBuy,
	)

	return err
}
