package repository_test

import (
	"context"
	"testing"

	"github.com/CzarSimon/httputil/testutil"
	"github.com/CzarSimon/tactics-trainer/iam-service/internal/models"
	"github.com/CzarSimon/tactics-trainer/iam-service/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestKeyEncryptionKeyRepositoryFind(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	db := testutil.InMemoryDB(true, "../../resources/db/sqlite")
	defer db.Close()

	k1, err := models.ParseKeyEncryptionKey("0:2557BC625F244FAFFDC817A7DDB6A20E09B4C1200C4A2B3ABAA49BA968476452")
	assert.NoError(err)
	k2, err := models.ParseKeyEncryptionKey("1:BCF39E65064252E1C4043A576ED9EA9AC84BCAFECEF9588D18479DC0F8A2AEA9")
	assert.NoError(err)

	keys := []models.KeyEncryptionKey{k1, k2}
	repo := repository.NewKeyEncryptionKeyRepository(keys, db)

	kek, err := repo.Find(ctx, 1)
	assert.NoError(err)
	assert.Equal(1, kek.ID)
	assert.NoError(kek.Valid())

	kek, err = repo.Find(ctx, 2)
	assert.Error(err)
	assert.Equal(0, kek.ID)
	assert.Error(kek.Valid())
}

func TestKeyEncryptionKeyRepositoryFindActive(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	db := testutil.InMemoryDB(true, "../../resources/db/sqlite")
	defer db.Close()

	k1, err := models.ParseKeyEncryptionKey("0:2557BC625F244FAFFDC817A7DDB6A20E09B4C1200C4A2B3ABAA49BA968476452")
	assert.NoError(err)
	k2, err := models.ParseKeyEncryptionKey("1:BCF39E65064252E1C4043A576ED9EA9AC84BCAFECEF9588D18479DC0F8A2AEA9")
	assert.NoError(err)

	keys := []models.KeyEncryptionKey{k1, k2}
	repo := repository.NewKeyEncryptionKeyRepository(keys, db)

	// Expected key 0 to be active due to inital db migration state.
	kek, err := repo.FindActive(ctx)
	assert.NoError(err)
	assert.Equal(0, kek.ID)
	assert.NoError(kek.Valid())

	emptyKek, err := repo.Find(ctx, 2)
	assert.Error(err)
	assert.Equal(0, emptyKek.ID)
	assert.Error(emptyKek.Valid())
}
