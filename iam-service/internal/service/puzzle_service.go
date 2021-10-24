package service

import (
	"context"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/tactics-trainer/iam-service/internal/models"
	"github.com/CzarSimon/tactics-trainer/iam-service/internal/repository"
	"github.com/opentracing/opentracing-go"
)

type PuzzleService struct {
	PuzzleRepo repository.PuzzleRepository
}

func (s *PuzzleService) GetPuzzle(ctx context.Context, id string) (models.Puzzle, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "puzzle_service_get_puzzle")
	defer span.Finish()

	puzzle, exists, err := s.PuzzleRepo.Find(ctx, id)
	if err != nil {
		return models.Puzzle{}, err
	}

	if !exists {
		return models.Puzzle{}, httputil.NotFoundError(err)
	}

	return puzzle, nil
}
