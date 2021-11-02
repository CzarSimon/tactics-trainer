package service

import (
	"context"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/tactics-trainer/puzzle-server/internal/models"
	"github.com/CzarSimon/tactics-trainer/puzzle-server/internal/repository"
	"github.com/opentracing/opentracing-go"
)

type ProblemSetService struct {
	ProblemSetRepo repository.ProblemSetRepository
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

func assertProblemSetAccess(set models.ProblemSet, userID string) error {
	if set.UserID != userID {
		return httputil.Forbiddenf("User(id=%s) does not have access to %s", userID, set)
	}

	return nil
}
