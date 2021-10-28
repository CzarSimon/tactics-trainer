package service

import (
	"context"
	"encoding/hex"
	"errors"
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

	err := s.assertNewUser(ctx, req.Username)
	if err != nil {
		return models.AuthenticationResponse{}, err
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

func (s *AuthenticationService) Login(ctx context.Context, req models.AuthenticationRequest) (models.AuthenticationResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authenication_service_login")
	defer span.Finish()

	user, err := s.findExistingUser(ctx, req.Username)
	if err != nil {
		return models.AuthenticationResponse{}, err
	}

	err = s.verifyPassword(ctx, user.Credentials, req.Password)
	if err != nil {
		return models.AuthenticationResponse{}, err
	}

	token, err := s.Issuer.Issue(user.JWTUser(), s.TokenLifetime)
	if err != nil {
		return models.AuthenticationResponse{}, httputil.InternalServerErrorf("failed to issue token: %w", err)
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

func (s *AuthenticationService) verifyPassword(ctx context.Context, credentials models.Credentials, password string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authenication_service_verify_password")
	defer span.Finish()

	ciphertext, salt, err := credentials.Decode()
	if err != nil {
		return err
	}

	hashtext, err := s.Cipher.Decrypt(ctx, ciphertext, credentials.DataEncryptionKey)
	if err != nil {
		return err
	}

	err = s.Hasher.Verify([]byte(password), salt, hashtext)
	if err != nil {
		if errors.Is(err, crypto.ErrHashMissmatch) {
			return httputil.Unauthorizedf("wrong password: %w", err)
		}

		return err
	}

	return nil
}

func (s *AuthenticationService) findExistingUser(ctx context.Context, username string) (models.User, error) {
	user, found, err := s.UserRepo.FindByUsername(ctx, username)
	if err != nil {
		return models.User{}, err
	}

	if !found {
		return models.User{}, httputil.Unauthorizedf("no user with username=%s found", username)
	}

	return user, nil
}

func (s *AuthenticationService) assertNewUser(ctx context.Context, username string) error {
	_, found, err := s.UserRepo.FindByUsername(ctx, username)
	if err != nil {
		return err
	}

	if found {
		return httputil.Conflictf("user with username=%s already exits", username)
	}

	return nil
}
