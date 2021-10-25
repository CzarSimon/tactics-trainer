package service

import (
	"context"

	"github.com/CzarSimon/httputil/crypto"
	"github.com/CzarSimon/tactics-trainer/iam-service/internal/models"
	"github.com/CzarSimon/tactics-trainer/iam-service/internal/repository"
)

// Cipher cryptosystem for generating sealed keys and encrypting and decrypting byte arrays.
type Cipher struct {
	KEKRepo repository.KeyEncryptionKeyRepository
}

// Encrypt generates DEK, enrcypts the plaintext and returns a sealed DEK.
func (s *Cipher) Encrypt(ctx context.Context, plaintext []byte) ([]byte, models.DataEncryptionKey, error) {
	key, err := crypto.GenerateAESKey()
	if err != nil {
		return nil, models.DataEncryptionKey{}, err
	}

	ciphertext, err := crypto.NewAESCipher(key).Encrypt(plaintext)
	if err != nil {
		return nil, models.DataEncryptionKey{}, err
	}

	kek, err := s.KEKRepo.FindActive(ctx)
	if err != nil {
		return nil, models.DataEncryptionKey{}, err
	}

	dek, err := kek.Encrypt(key)
	if err != nil {
		return nil, models.DataEncryptionKey{}, err
	}

	return ciphertext, dek, nil
}

// Decrypt decrypts DEK using the referenced KEK and finally decrypts the ciphertext.
func (s *Cipher) Decrypt(ctx context.Context, ciphertext []byte, dek models.DataEncryptionKey) ([]byte, error) {
	kek, err := s.KEKRepo.Find(ctx, dek.KeyEncryptionKeyID)
	if err != nil {
		return nil, err
	}

	key, err := kek.Decrypt(dek)
	if err != nil {
		return nil, err
	}

	plaintext, err := crypto.NewAESCipher(key).Decrypt(ciphertext)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
