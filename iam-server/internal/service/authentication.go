package service

import (
	"context"
	"encoding/hex"
	"time"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/httputil/crypto"
	"github.com/CzarSimon/httputil/jwt"
	"github.com/CzarSimon/tactics-trainer/iam-server/internal/models"
	"github.com/CzarSimon/tactics-trainer/iam-server/internal/repository"
	"github.com/opentracing/opentracing-go"
)

// AuthenticationService business logic for handling authentication.
type AuthenticationService struct {
	Issuer        jwt.Issuer
	Cipher        *Cipher
	Hasher        crypto.Hasher
	UserRepo      repository.UserRepository
	TokenLifetime time.Duration
}

func (s *AuthenticationService) Signup(ctx context.Context, req models.AuthenticationRequest) (models.AuthenticationResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authenication_service_signup")
	defer span.Finish()

	_, found, err := s.UserRepo.FindByUsername(ctx, req.Username)
	if err != nil {
		return models.AuthenticationResponse{}, err
	}

	if found {
		return models.AuthenticationResponse{}, httputil.Conflictf("user with username=%s already exits", req.Username)
	}

	creds, err := s.hashPassword(ctx, req.Password)
	if err != nil {
		return models.AuthenticationResponse{}, httputil.InternalServerErrorf("failed to hash password: %w", err)
	}

	user := models.NewUser(req.Username, models.UserRole, creds)
	token, err := s.Issuer.Issue(user.JWTUser(), s.TokenLifetime)
	if err != nil {
		return models.AuthenticationResponse{}, httputil.InternalServerErrorf("failed to issue token: %w", err)
	}

	err = s.UserRepo.Save(ctx, user)
	if err != nil {
		return models.AuthenticationResponse{}, err
	}

	return models.AuthenticationResponse{
		Token: token,
		User:  user,
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
