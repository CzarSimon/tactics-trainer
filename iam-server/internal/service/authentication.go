package service

import (
	"context"
	"encoding/hex"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/httputil/crypto"
	"github.com/CzarSimon/httputil/jwt"
	"github.com/CzarSimon/tactics-trainer/iam-server/internal/models"
	"github.com/CzarSimon/tactics-trainer/iam-server/internal/repository"
	"github.com/opentracing/opentracing-go"
)

// AuthenticationService business logic for handling authentication.
type AuthenticationService struct {
	Issuer   jwt.Issuer
	Cipher   *Cipher
	Hasher   crypto.Hasher
	UserRepo repository.UserRepository
}

func (s *AuthenticationService) Signup(ctx context.Context, req models.AuthenticationRequest) (models.AuthenticationResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authenication_service_signup")
	defer span.Finish()

	creds, err := s.hashPassword(ctx, req.Password)
	if err != nil {
		return models.AuthenticationResponse{}, httputil.InternalServerErrorf("failed to hash password: %w", err)
	}

	user := models.NewUser(req.Username, models.UserRole, creds)
	err = s.UserRepo.Save(ctx, user)
	if err != nil {
		return models.AuthenticationResponse{}, err
	}

	return models.AuthenticationResponse{
		User: user,
	}, nil
}

func (s *AuthenticationService) hashPassword(ctx context.Context, password string) (models.Credentials, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authenication_service_hash_password")
	defer span.Finish()

	salt, err := crypto.GenerateAESKey()
	if err != nil {
		return models.Credentials{}, err
	}

	hash, err := s.Hasher.Hash([]byte(password), salt)
	if err != nil {
		return models.Credentials{}, err
	}

	ciphertext, dek, err := s.Cipher.Encrypt(ctx, hash)
	if err != nil {
		return models.Credentials{}, err
	}

	return models.Credentials{
		Password:          hex.EncodeToString(ciphertext),
		Salt:              hex.EncodeToString(salt),
		DataEncryptionKey: dek,
	}, nil
}
