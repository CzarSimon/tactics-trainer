package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/httputil/timeutil"
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

	_, err = s.getCycleProblemSet(ctx, cycle, userID)
	if err != nil {
		return models.Cycle{}, err
	}

	return cycle, nil
}

func (s *CycleService) UpdateCycle(ctx context.Context, id, userID string) (models.Cycle, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "cycle_service_update_cycle")
	defer span.Finish()

	cycle, err := s.GetCycle(ctx, id, userID)
	if err != nil {
		return models.Cycle{}, err
	}

	set, err := s.getCycleProblemSet(ctx, cycle, userID)
	if err != nil {
		return models.Cycle{}, err
	}

	if cycle.Compleated() {
		err = httputil.Errorf(http.StatusUnprocessableEntity, "%s is already compleated", cycle)
		return models.Cycle{}, err
	}

	cycle.UpdatedAt = timeutil.Now()
	nextID := getNextPuzzleID(set.PuzzleIDs, cycle.CurrentPuzzleID)
	if nextID != "" {
		cycle.CurrentPuzzleID = nextID
	} else {
		cycle.CompleatedAt = cycle.UpdatedAt
	}

	err = s.CycleRepo.Update(ctx, cycle)
	if err != nil {
		return models.Cycle{}, err
	}

	return cycle, nil
}

func (s *CycleService) getCycleProblemSet(ctx context.Context, cycle models.Cycle, userID string) (models.ProblemSet, error) {
	set, found, err := s.ProblemSetRepo.Find(ctx, cycle.ProblemSetID)
	if err != nil {
		return models.ProblemSet{}, err
	}
	if !found {
		return models.ProblemSet{}, fmt.Errorf("%s refers to a problem set that does not exist", cycle)
	}

	if set.UserID != userID {
		return models.ProblemSet{}, httputil.Forbiddenf("User(id=%s) is not the owner of %s", userID, set)
	}

	return set, nil
}

func getNextPuzzleID(puzzleIDs []string, currentID string) string {
	numberOfPuzzles := len(puzzleIDs)
	for i, id := range puzzleIDs {
		if id != currentID {
			continue
		}

		if i+1 < numberOfPuzzles {
			return puzzleIDs[i+1]
		}
	}

	return ""
}
