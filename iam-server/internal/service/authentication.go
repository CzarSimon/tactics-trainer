package service

import (
	"context"

	"github.com/CzarSimon/tactics-trainer/iam-server/internal/models"
	"github.com/CzarSimon/tactics-trainer/iam-server/internal/repository"
)

// AuthenticationService business logic for handling authentication.
type AuthenticationService struct {
	Cipher   *Cipher
	UserRepo repository.UserRepository
}

func (s *AuthenticationService) Signup(ctx context.Context, req models.AuthenticationRequest) (models.AuthenticationResponse, error) {
	return models.AuthenticationResponse{}, nil
}
