package inbound

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Repository interface {
	CreateInbound(ctx context.Context, tx pgx.Tx, inbound *Inbound) (int, error)
	CreateInboundItem(ctx context.Context, tx pgx.Tx, inboundID int, item InboundItem) error
}
