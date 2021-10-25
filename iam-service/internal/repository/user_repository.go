package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/CzarSimon/tactics-trainer/iam-service/internal/models"
	"github.com/opentracing/opentracing-go"
)

type UserRepository interface {
	Save(ctx context.Context, user models.User) error
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepo{
		db: db,
	}
}

type userRepo struct {
	db *sql.DB
}

const saveUserQuery = `
	INSERT INTO user_account(id, username, role, password, salt, data_encryption_key, key_encryption_key_id, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
`

func (r *userRepo) Save(ctx context.Context, user models.User) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "user_repo_save")
	defer span.Finish()

	dek := user.DataEncryptionKey
	_, err := r.db.Exec(saveUserQuery, user.ID, user.Username, user.Role, user.Password, user.Salt, dek.Body, dek.KeyEncryptionKeyID, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to save %s: %w", user, err)
	}

	return nil
}
