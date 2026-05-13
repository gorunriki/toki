package postgres

import (
	"context"

	"toki/internal/domain/user"

	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) user.Repository {
	return &userRepository{db}
}

func (r *userRepository) FindByEmail(
	ctx context.Context,
	email string,
) (*user.User, error) {

	query := `
	SELECT id, name, email, password, role
	FROM users
	WHERE email=$1
	`

	row := r.db.QueryRow(ctx, query, email)

	var u user.User

	err := row.Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.Password,
		&u.Role,
	)

	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *userRepository) Create(
	ctx context.Context,
	u *user.User,
) error {

	query := `
	INSERT INTO users (
		name,
		email,
		password,
		role
	)
	VALUES ($1,$2,$3,$4)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		u.Name,
		u.Email,
		u.Password,
		u.Role,
	)

	return err
}
