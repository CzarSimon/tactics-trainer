package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/CzarSimon/tactics-trainer/iam-service/internal/models"
)

// KeyEncryptionKeyRepository interface for KeyEncryptionKey (KEK) data access layer
type KeyEncryptionKeyRepository interface {
	Find(context.Context, int) (models.KeyEncryptionKey, error)
	FindActive(context.Context) (models.KeyEncryptionKey, error)
}

// NewKeyEncryptionKeyRepository creates an KeyEncryptionKeyRepository using the default implementation.
func NewKeyEncryptionKeyRepository(keys []models.KeyEncryptionKey, db *sql.DB) KeyEncryptionKeyRepository {
	keyMap := make(map[int]models.KeyEncryptionKey)
	for _, kek := range keys {
		keyMap[kek.ID] = kek
	}

	return &kekRepo{
		keys: keyMap,
		db:   db,
	}
}

type kekRepo struct {
	// Writing to the keys map is only done at the time that the kekRepo
	// is initialized and not supported at runtime, thus a mutex is not needed
	// to protect read access from multilple gorutines.
	keys map[int]models.KeyEncryptionKey
	db   *sql.DB
}

func (r *kekRepo) Find(ctx context.Context, id int) (models.KeyEncryptionKey, error) {
	kek, ok := r.keys[id]
	if !ok {
		return kek, fmt.Errorf("could not find KeyEncryptionKey(id=%d)", id)
	}

	return kek, nil
}

func (r *kekRepo) FindActive(ctx context.Context) (models.KeyEncryptionKey, error) {
	id, err := r.findActiveKeyEncryptionKeyID(ctx)
	if err != nil {
		return models.KeyEncryptionKey{}, err
	}

	return r.Find(ctx, id)
}

const findLatestKeyEncryptionKeyByStateQuery = `
	SELECT id FROM key_encryption_key WHERE state = ? OR state = ? ORDER BY created_at DESC LIMIT 1`

func (r *kekRepo) findActiveKeyEncryptionKeyID(ctx context.Context) (int, error) {
	var id int
	err := r.db.QueryRowContext(ctx, findLatestKeyEncryptionKeyByStateQuery, models.KeyStateActive, models.KeyStateNext).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to query for current key_encryption_key. Error: %w", err)
	}

	return id, nil
}
