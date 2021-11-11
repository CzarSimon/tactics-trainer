package service

import (
	"context"
	"fmt"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/tactics-trainer/puzzle-server/internal/models"
	"github.com/CzarSimon/tactics-trainer/puzzle-server/internal/repository"
	"github.com/opentracing/opentracing-go"
)

type CycleService struct {
	CycleRepo      repository.CycleRepository
	ProblemSetRepo repository.ProblemSetRepository
}

func (s *CycleService) GetCycle(ctx context.Context, id, userID string) (models.Cycle, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "cycle_service_get_cycle")
	defer span.Finish()

	cycle, found, err := s.CycleRepo.Find(ctx, id)
	if err != nil {
		return models.Cycle{}, err
	}
	if !found {
		return models.Cycle{}, httputil.NotFoundError(err)
	}

	set, found, err := s.ProblemSetRepo.Find(ctx, cycle.ProblemSetID)
	if err != nil {
		return models.Cycle{}, err
	}
	if !found {
		return models.Cycle{}, fmt.Errorf("%s refers to a problem set that does not exist", cycle)
	}

	if set.UserID != userID {
		return models.Cycle{}, httputil.Forbiddenf("User(id=%s) is not the owner of %s", userID, set)
	}

	return cycle, nil
}
