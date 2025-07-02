package repository

import (
	"database/sql"
	"rozoomcool/go-api-skeleton/internal/models"
	// "github.com/ydb-platform/ydb-go-sdk/v3/query"
)

type UserRepo interface {
	GetById(id string) (*models.User, error)
}

type pgUserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepo {
	return &pgUserRepo{db}
}

// GetById implements UserRepo.
func (r *pgUserRepo) GetById(id string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT guid FROM users WHERE id = $1`
	if err := r.db.QueryRow(query, id).Scan(&user.GUID); err != nil {
		return nil, err
	}

	return user, nil
}
