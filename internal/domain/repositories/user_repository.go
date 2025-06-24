package repositories

import (
	"errors"
	"database/sql"
	"sewerage/internal/domain/entities"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repository *UserRepository) GetByEmail(emial string) (*entities.User, error) {
	return nil, errors.New("empty name")
}

// func (r *UserRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {
// 	err := r.db.QueryRowContext(ctx,
// 		"INSERT INTO users (...) VALUES (...) RETURNING id",
// 		user.Email,
// 	).Scan(&user.ID)

// 	return user, err
// }