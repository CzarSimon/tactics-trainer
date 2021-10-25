package repository_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/CzarSimon/httputil/id"
	"github.com/CzarSimon/httputil/testutil"
	"github.com/CzarSimon/httputil/timeutil"
	"github.com/CzarSimon/tactics-trainer/iam-service/internal/models"
	"github.com/CzarSimon/tactics-trainer/iam-service/internal/repository"
	"github.com/stretchr/testify/assert"
)

func Test_userRepo_Save(t *testing.T) {
	assert := assert.New(t)
	db := testutil.InMemoryDB(true, "../../resources/db/sqlite")
	repo := repository.NewUserRepository(db)
	ctx := context.Background()

	user := models.User{
		ID:       id.New(),
		Username: "test-user",
		Role:     models.UserRole,
		Password: "test-password",
		Salt:     "some-salt",
		DataEncryptionKey: models.DataEncryptionKey{
			Body:               "hex-encoded-string",
			KeyEncryptionKeyID: 0,
		},
		CreatedAt: timeutil.Now(),
		UpdatedAt: timeutil.Now(),
	}

	var foundID string
	err := db.QueryRow("SELECT id FROM user_account WHERE id = ?", user.ID).Scan(&foundID)
	assert.Equal(sql.ErrNoRows, err)
	assert.Empty(foundID)

	err = repo.Save(ctx, user)
	assert.NoError(err)

	err = db.QueryRow("SELECT id FROM user_account WHERE id = ?", user.ID).Scan(&foundID)
	assert.NoError(err)
	assert.Equal(user.ID, foundID)
}
