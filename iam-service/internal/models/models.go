package models

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/CzarSimon/httputil/crypto"
)

var errIncorrectKeyString = errors.New("incorrectly formated key encryption key string")

// User account describing a user of the object storage service.
type User struct {
	ID                string
	Username          string
	Password          string
	Salt              string
	DataEncryptionKey DataEncryptionKey
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (u User) String() string {
	return fmt.Sprintf("User(id=%s, username=%s, createdAt=%v, updatedAt=%v)", u.ID, u.Username, u.CreatedAt, u.UpdatedAt)
}

// DataEncryptionKey (DEK) key to encrypt and decrypt a objects.
// A DEK is encrypted with a KeyEncryptionKey.
type DataEncryptionKey struct {
	Body               string // Body is a hex encoded string of an encrypted byte array.
	KeyEncryptionKeyID int
}

func (d DataEncryptionKey) String() string {
	return fmt.Sprintf("DataEncryptionKey(keyEncryptionKeyID=%d)", d.KeyEncryptionKeyID)
}

// KeyEncryptionKey (KEK) key to encrypt and decrypt a data encryption key.
type KeyEncryptionKey struct {
	ID   int
	body []byte // Body is unexported to prevent accidental logging of it.
}

// Valid checks that a key is valid by asserting its size.
func (k KeyEncryptionKey) Valid() error {
	if len(k.body) != crypto.AES256KeySize {
		return fmt.Errorf("invalid key size of %s. must be: %d. got: %d", k, crypto.AES256KeySize, len(k.body))
	}

	return nil
}

// Encrypt encrypts the provided data encryption key with with the key encryption key body.
// Encodes the result as a hex string.
func (k KeyEncryptionKey) Encrypt(plaintext []byte) (DataEncryptionKey, error) {
	ciphertext, err := crypto.NewAESCipher(k.body).Encrypt(plaintext)
	if err != nil {
		return DataEncryptionKey{}, fmt.Errorf("failed to encrypt DEK using %s. Error: %w", k, err)
	}

	dek := DataEncryptionKey{
		Body:               hex.EncodeToString(ciphertext),
		KeyEncryptionKeyID: k.ID,
	}
	return dek, nil
}

// Decrypt decodes the hex encoded ciphertext and decrypts the encoded
// data encryption key with with the key encryption key body.
func (k KeyEncryptionKey) Decrypt(dek DataEncryptionKey) ([]byte, error) {
	if k.ID != dek.KeyEncryptionKeyID {
		return nil, fmt.Errorf("wrong KeyEncryptionKey. cannot use %s to decrypt %s", k, dek)
	}

	ciphertext, err := hex.DecodeString(dek.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to decode DEK using %s. Error: %w", k, err)
	}

	plaintext, err := crypto.NewAESCipher(k.body).Decrypt(ciphertext)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt DEK using %s. Error: %w", k, err)
	}

	return plaintext, nil
}

func (k KeyEncryptionKey) String() string {
	return fmt.Sprintf("KeyEncryptionKey(id=%d)", k.ID)
}

// ParseKeyEncryptionKey parses a hex encode key encryption key and id.
func ParseKeyEncryptionKey(s string) (KeyEncryptionKey, error) {
	parts := strings.Split(s, ":")
	if len(parts) != 2 {
		return KeyEncryptionKey{}, fmt.Errorf("%w. expected 2 parts got %d", errIncorrectKeyString, len(parts))
	}

	id, err := strconv.Atoi(parts[0])
	if err != nil {
		return KeyEncryptionKey{}, fmt.Errorf("%w. expected id to be integer", errIncorrectKeyString)
	}

	body, err := hex.DecodeString(parts[1])
	if err != nil {
		return KeyEncryptionKey{}, fmt.Errorf("%w. expected hex encoded body", errIncorrectKeyString)
	}

	if len(body) != crypto.AES256KeySize {
		return KeyEncryptionKey{}, fmt.Errorf("%w. expected key length %d got %d", errIncorrectKeyString, crypto.AES256KeySize, len(body))
	}

	return KeyEncryptionKey{
		ID:   id,
		body: body,
	}, nil
}
