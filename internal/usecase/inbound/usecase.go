package inbound

import (
	"context"

	domainInbound "toki/internal/domain/inbound"
	"toki/internal/domain/stock"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Usecase interface {
	Create(
		ctx context.Context,
		req *domainInbound.Inbound,
	) error
}

type usecase struct {
	repo      domainInbound.Repository
	stockRepo stock.Repository
	db        *pgxpool.Pool
}

func NewUsecase(
	repo domainInbound.Repository,
	stockRepo stock.Repository,
	db *pgxpool.Pool,
) Usecase {

	return &usecase{
		repo:      repo,
		stockRepo: stockRepo,
		db:        db,
	}
}

func (u *usecase) Create(
	ctx context.Context,
	req *domainInbound.Inbound,
) error {

	// begin transaction
	tx, err := u.db.Begin(ctx)

	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	// create inbound
	inboundID, err := u.repo.CreateInbound(
		ctx,
		tx,
		req,
	)

	if err != nil {
		return err
	}

	// create inbound items
	for _, item := range req.Items {

		err = u.repo.CreateInboundItem(
			ctx,
			tx,
			inboundID,
			item,
		)

		if err != nil {
			return err
		}

		// add stock
		err = u.stockRepo.AddStock(
			ctx,
			item.ItemID,
			item.Qty,
		)

		if err != nil {
			return err
		}
	}

	// commit transaction
	err = tx.Commit(ctx)

	if err != nil {
		return err
	}

	return nil
}
