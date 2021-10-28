package service_test

import (
	"context"
	"encoding/hex"
	"testing"

	"github.com/CzarSimon/httputil/crypto"
	"github.com/CzarSimon/httputil/testutil"
	"github.com/CzarSimon/tactics-trainer/iam-server/internal/models"
	"github.com/CzarSimon/tactics-trainer/iam-server/internal/repository"
	"github.com/CzarSimon/tactics-trainer/iam-server/internal/service"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestCipherService(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	db := testutil.InMemoryDB(true, "../../resources/db/sqlite")
	defer db.Close()

	k1, err := models.ParseKeyEncryptionKey("0:2557BC625F244FAFFDC817A7DDB6A20E09B4C1200C4A2B3ABAA49BA968476452")
	assert.NoError(err)
	k2, err := models.ParseKeyEncryptionKey("1:BCF39E65064252E1C4043A576ED9EA9AC84BCAFECEF9588D18479DC0F8A2AEA9")
	assert.NoError(err)

	keys := []models.KeyEncryptionKey{k1, k2}
	svc := service.Cipher{
		KEKRepo: repository.NewKeyEncryptionKeyRepository(keys, db),
	}

	data := []byte("some data that should be encrypted")

	ciphertext, dek, err := svc.Encrypt(ctx, data)
	assert.NoError(err)
	assert.NotEqual(string(data), string(ciphertext))
	assert.Equal(0, dek.KeyEncryptionKeyID)

	encryptedKey, err := hex.DecodeString(dek.Body)
	assert.NoError(err)

	_, err = crypto.NewAESCipher(encryptedKey).Decrypt(ciphertext)
	assert.Error(err)

	plaintext, err := svc.Decrypt(ctx, ciphertext, dek)
	assert.NoError(err)
	assert.Equal(string(data), string(plaintext))
}
