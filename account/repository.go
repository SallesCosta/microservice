package account

import (
	"context"
	"database/sql"
)

type Repository interface {
	Close()
	PutAccount(ctx context.Context, a Account) error
	GetAccountByID(ctx context.Context, id string) (*Account, error)
	ListAccounts(ctx context.Context, skip uint64) ([]Account, error)
}

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (Repository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &postgresRepository{db}, nil
}

func (r *postgresRepository) Close() {
	_ = r.db.Close()
}

func (r *postgresRepository) Pint() error {
	_ = r.db.Ping()
}

func (r *postgresRepository) PutAccount(ctx context.Context, a Account) error {

	return nil
}

func (r *postgresRepository) GetAccountByID(ctx context.Context, id string) (*Account, error) {
	return nil, nil
}

func (r *postgresRepository) ListAccounts(ctx context.Context, skip uint64) ([]Account, error) {
	return nil, nil
}
