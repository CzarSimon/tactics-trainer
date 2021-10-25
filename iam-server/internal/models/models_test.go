package models_test

import (
	"encoding/hex"
	"testing"

	"github.com/CzarSimon/httputil/crypto"
	"github.com/CzarSimon/tactics-trainer/iam-server/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestParseKeyEncryptionKey(t *testing.T) {
	assert := assert.New(t)

	kek, err := models.ParseKeyEncryptionKey("8:0D17FE0FDA5F1CE46307561714C6938FDFE9408BF6712BE49D3FC4C757D0E62E")
	assert.NoError(err)
	assert.Equal(8, kek.ID)

	_, err = models.ParseKeyEncryptionKey("8:-D1%F0FD<A5F1<E46307561714C6938FDFE9408BF6712BE49D3FC4C757D0E62>")
	assert.Error(err)

	_, err = models.ParseKeyEncryptionKey("X:80A74BE621F75875BE56E4CD1AE1B2A8DF4722FB6C6A107779DDB198C78E4DAC")
	assert.Error(err)

	_, err = models.ParseKeyEncryptionKey("8:80A74BE621F75875BE56E4CD1AE1B2A8")
	assert.Error(err)
}

func TestEncryptAndDecryptKey(t *testing.T) {
	assert := assert.New(t)
	kek, err := models.ParseKeyEncryptionKey("40:E68256A8F9685DA6BE7C5FE461C0838E0751ED547E9A824B2BDA9E2580703643")
	assert.NoError(err)

	otherKek, err := models.ParseKeyEncryptionKey("12:E68256A8F9685DA6BE7C5FE461C0838E0751ED547E9A824B2BDA9E2580703643")
	assert.NoError(err)

	wrongKek, err := models.ParseKeyEncryptionKey("40:0D17FE0FDA5F1CE46307561714C6938FDFE9408BF6712BE49D3FC4C757D0E62E")
	assert.NoError(err)

	key, err := crypto.GenerateAESKey()
	assert.NoError(err)

	dek, err := kek.Encrypt(key)
	assert.NoError(err)
	assert.Equal(kek.ID, dek.KeyEncryptionKeyID)

	_, err = otherKek.Decrypt(dek)
	assert.Error(err)

	_, err = wrongKek.Decrypt(dek)
	assert.Error(err)

	plaintext, err := kek.Decrypt(dek)
	assert.NoError(err)

	assert.Equal(hex.EncodeToString(key), hex.EncodeToString(plaintext))
}
