package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/httputil/id"
	"github.com/CzarSimon/httputil/timeutil"
	"github.com/CzarSimon/tactics-trainer/puzzle-server/internal/models"
	"github.com/CzarSimon/tactics-trainer/puzzle-server/internal/repository"
	"github.com/opentracing/opentracing-go"
)

type ProblemSetService struct {
	ProblemSetRepo repository.ProblemSetRepository
	PuzzleRepo     repository.PuzzleRepository
	CycleRepo      repository.CycleRepository
}

func (s *ProblemSetService) ListProblemSets(ctx context.Context, userID string, includeArchived bool) ([]models.ProblemSet, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "problem_set_service_list_problem_sets")
	defer span.Finish()

	sets, err := s.ProblemSetRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if !includeArchived {
		sets = filterArchivedSets(sets)
	}

	return sets, nil
}

func (s *ProblemSetService) CreateProblemSet(ctx context.Context, req models.CreateProblemSetRequest, userID string) (models.ProblemSet, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "problem_set_service_create_problem_set")
	defer span.Finish()

	puzzleIDs, err := s.getPuzzleIDs(ctx, req.Filter)
	if err != nil {
		return models.ProblemSet{}, err
	} else if len(puzzleIDs) < 1 {
		return models.ProblemSet{}, httputil.Errorf(http.StatusUnprocessableEntity, "%s returned no puzzles", req.Filter)
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

func (s *ProblemSetService) ArchiveProblemSet(ctx context.Context, id, userID string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "problem_set_service_archive_problem_set")
	defer span.Finish()

	set, err := s.GetProblemSet(ctx, id, userID)
	if err != nil {
		return err
	}

	set.Archived = true
	err = s.ProblemSetRepo.Update(ctx, set)
	if err != nil {
		return err
	}

	return nil
}

func (s *ProblemSetService) CreateProblemSetCycle(ctx context.Context, id, userID string) (models.Cycle, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "problem_set_service_create_problem_set_cycle")
	defer span.Finish()

	set, err := s.GetProblemSet(ctx, id, userID)
	if err != nil {
		return models.Cycle{}, err
	}

	cycle, err := s.createNewCycle(ctx, set)
	if err != nil {
		return models.Cycle{}, err
	}

	err = s.CycleRepo.Save(ctx, cycle)
	if err != nil {
		return models.Cycle{}, err
	}

	return cycle, nil
}

func (s *ProblemSetService) ListProblemSetCycles(ctx context.Context, id, userID string, onlyActive bool) ([]models.Cycle, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "problem_set_service_list_problem_set_cycles")
	defer span.Finish()

	set, err := s.GetProblemSet(ctx, id, userID)
	if err != nil {
		return nil, err
	}

	cycles, err := s.CycleRepo.FindByProblemSetID(ctx, set.ID, onlyActive)
	if err != nil {
		return nil, err
	}

	return cycles, nil
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

func (s *ProblemSetService) createNewCycle(ctx context.Context, set models.ProblemSet) (models.Cycle, error) {
	if len(set.PuzzleIDs) < 1 {
		return models.Cycle{}, fmt.Errorf("%s contained no puzzles", set)
	}

	cycles, err := s.CycleRepo.FindByProblemSetID(ctx, set.ID, false)
	if err != nil {
		return models.Cycle{}, err
	}

	now := timeutil.Now()
	cycle := models.Cycle{
		ID:              id.New(),
		Number:          len(cycles) + 1,
		ProblemSetID:    set.ID,
		CurrentPuzzleID: set.PuzzleIDs[0],
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	return cycle, nil
}

func assertProblemSetAccess(set models.ProblemSet, userID string) error {
	if set.UserID != userID {
		return httputil.Forbiddenf("User(id=%s) does not have access to %s", userID, set)
	}

	return nil
}

func filterArchivedSets(sets []models.ProblemSet) []models.ProblemSet {
	filteredSets := make([]models.ProblemSet, 0)
	for _, set := range sets {
		if !set.Archived {
			filteredSets = append(filteredSets, set)
		}
	}

	return filteredSets
}
