package service

import (
	"context"
	"net/http"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/tactics-trainer/puzzle-server/internal/models"
	"github.com/CzarSimon/tactics-trainer/puzzle-server/internal/repository"
	"github.com/opentracing/opentracing-go"
)

type ProblemSetService struct {
	ProblemSetRepo repository.ProblemSetRepository
	PuzzleRepo     repository.PuzzleRepository
}

func (s *ProblemSetService) CreateProblemSet(ctx context.Context, req models.CreateProblemSetRequest, userID string) (models.ProblemSet, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "problem_set_service_create_problem_set")
	defer span.Finish()

	puzzleIDs, err := s.getPuzzleIDs(ctx, req.Filter)
	if err != nil {
		return models.ProblemSet{}, err
	} else if len(puzzleIDs) < 1 {
		return models.ProblemSet{}, httputil.Errorf(http.StatusUnprocessableEntity, "the filter returned no puzzles", req.Filter)
	}

	set := models.NewProblemSet(req, userID, puzzleIDs)
	err = s.ProblemSetRepo.Save(ctx, set)
	if err != nil {
		return models.ProblemSet{}, err
	}

	return set, nil
}

func (s *ProblemSetService) GetProblemSet(ctx context.Context, id, userID string) (models.ProblemSet, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "problem_set_service_get_problem_set")
	defer span.Finish()

	set, found, err := s.ProblemSetRepo.Find(ctx, id)
	if err != nil {
		return models.ProblemSet{}, err
	} else if !found {
		return models.ProblemSet{}, httputil.NotFoundf("no problem set found with id=%s", id)
	}

	err = assertProblemSetAccess(set, userID)
	if err != nil {
		return models.ProblemSet{}, err
	}

	return set, nil
}

func (s *ProblemSetService) getPuzzleIDs(ctx context.Context, f models.PuzzleFilter) ([]string, error) {
	puzzles, err := s.PuzzleRepo.FindByFilter(ctx, f)
	if err != nil {
		return nil, err
	}

	ids := make([]string, len(puzzles))
	for i, puzzle := range puzzles {
		ids[i] = puzzle.ID
	}

	return ids, nil
}

func assertProblemSetAccess(set models.ProblemSet, userID string) error {
	if set.UserID != userID {
		return httputil.Forbiddenf("User(id=%s) does not have access to %s", userID, set)
	}

	return nil
}
