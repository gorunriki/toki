package sales

import (
	"context"
	"errors"

	domain "toki/internal/domain/sales"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Usecase interface {
	Create(ctx context.Context, sales *domain.Sales) error
}

type salesUsecase struct {
	repo domain.Repository
	db   *pgxpool.Pool
}

func NewUsecase(
	r domain.Repository,
	db *pgxpool.Pool,
) Usecase {
	return &salesUsecase{
		repo: r,
		db:   db,
	}
}

func (u *salesUsecase) Create(
	ctx context.Context,
	s *domain.Sales,
) error {

	tx, err := u.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	salesID, err := u.repo.CreateSales(ctx, tx, s)
	if err != nil {
		return err
	}

	for _, item := range s.Items {

		var currentStock int

		err := tx.QueryRow(ctx,
			`SELECT quantity FROM stocks WHERE item_id=$1`,
			item.ItemID,
		).Scan(&currentStock)

		if err != nil {
			return err
		}

		if currentStock < item.Qty {
			return errors.New("stock not enough")
		}

		err = u.repo.CreateSalesItem(
			ctx,
			tx,
			salesID,
			item,
		)

		if err != nil {
			return err
		}

		_, err = tx.Exec(ctx,
			`
			UPDATE stocks
			SET quantity = quantity - $1
			WHERE item_id = $2
			`,
			item.Qty,
			item.ItemID,
		)

		if err != nil {
			return err
		}

		_, err = tx.Exec(ctx,
			`
			INSERT INTO stock_movements (
				item_id,
				type,
				qty,
				reference_type,
				reference_id
			)
			VALUES ($1,'OUT',$2,'SALES',$3)
			`,
			item.ItemID,
			item.Qty,
			salesID,
		)

		if err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}
