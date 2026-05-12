package inbound

import (
	"context"

	domain "toki/internal/domain/inbound"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Usecase interface {
	Create(ctx context.Context, inbound *domain.Inbound) error
}

type inboundUsecase struct {
	repo domain.Repository
	db   *pgxpool.Pool
}

func NewUsecase(
	r domain.Repository,
	db *pgxpool.Pool,
) Usecase {
	return &inboundUsecase{
		repo: r,
		db:   db,
	}
}

func (u *inboundUsecase) Create(
	ctx context.Context,
	in *domain.Inbound,
) error {

	tx, err := u.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	inboundID, err := u.repo.CreateInbound(ctx, tx, in)
	if err != nil {
		return err
	}

	for _, item := range in.Items {

		err := u.repo.CreateInboundItem(
			ctx,
			tx,
			inboundID,
			item,
		)
		if err != nil {
			return err
		}

		_, err = tx.Exec(ctx, `
			UPDATE stocks
			SET quantity = quantity + $1
			WHERE item_id = $2
		`,
			item.Qty,
			item.ItemID,
		)

		if err != nil {
			return err
		}

		_, err = tx.Exec(ctx, `
			INSERT INTO stock_movements (
				item_id,
				type,
				qty,
				reference_type,
				reference_id
			)
			VALUES ($1,'IN',$2,'INBOUND',$3)
		`,
			item.ItemID,
			item.Qty,
			inboundID,
		)

		if err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}
