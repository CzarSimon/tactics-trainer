package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/CzarSimon/httputil/id"
	"github.com/CzarSimon/httputil/timeutil"
)

// ProblemSet represents a collection of chess puzzles
type ProblemSet struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description,omitempty"`
	Themes         []string  `json:"themes"`
	RatingInterval string    `json:"ratingInterval"`
	UserID         string    `json:"userId"`
	Archived       bool      `json:"archived"`
	PuzzleIDs      []string  `json:"puzzleIds"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

func (p ProblemSet) String() string {
	return fmt.Sprintf(
		"ProblemSet(id=%s, name=%s, ratingInterval=%s, userId=%s)",
		p.ID,
		p.Name,
		p.RatingInterval,
		p.UserID,
	)
}

func NewProblemSet(req CreateProblemSetRequest, userID string, puzzleIDs []string) ProblemSet {
	now := timeutil.Now()
	return ProblemSet{
		ID:             id.New(),
		Name:           req.Name,
		Description:    req.Description,
		Themes:         req.Filter.Themes,
		RatingInterval: req.Filter.RatingInterval(),
		UserID:         userID,
		Archived:       false,
		PuzzleIDs:      puzzleIDs,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
}

// ProblemSetPuzzle creates a ProblemSetPuzzle struct based on a puzzle id in the
// underlying problem sets PuzzleIDs array. Please note that the number passed to
// the method MUST be within the index range of ProblemSets.PuzzleIDs slice.
// If number is outside the index range of the PuzzleIDs slice this method will panic.
func (p ProblemSet) ProblemSetPuzzle(number int) ProblemSetPuzzle {
	return ProblemSetPuzzle{
		ID:           id.New(),
		PuzzleID:     p.PuzzleIDs[number], // Attention! This will cause a nilPointerException if number is outside the index range of p.PuzzleIDs
		ProblemSetID: p.ID,
		Number:       number,
	}
}

// ProblemSetPuzzle mapping struct between a problem set and a puzzle.
type ProblemSetPuzzle struct {
	ID           string
	PuzzleID     string
	ProblemSetID string
	Number       int
}

func (p ProblemSetPuzzle) String() string {
	return fmt.Sprintf(
		"ProblemSetPuzzle(id=%s, puzzleId=%s, problemSetId=%s, number=%d)",
		p.ID,
		p.PuzzleID,
		p.ProblemSetID,
		p.Number,
	)
}

// CreateProblemSetRequest request to create a problem set
type CreateProblemSetRequest struct {
	Name        string       `json:"name"`
	Description string       `json:"description,omitempty"`
	Filter      PuzzleFilter `json:"filter"`
}

func (r CreateProblemSetRequest) String() string {
	return fmt.Sprintf("CreateProblemSetRequest(name=%s, filter=%s)", r.Name, r.Filter)
}

func (r CreateProblemSetRequest) Valid() error {
	if r.Name == "" {
		return errors.New("name must be set")
	}

	return r.Filter.Valid()
}

// PuzzleFilter selection criteria of puzzles.
type PuzzleFilter struct {
	Themes        []string `json:"themes"`
	MinRating     uint     `json:"minRating,omitempty"`
	MaxRating     uint     `json:"maxRating,omitempty"`
	MinPopularity uint     `json:"minPopularity"`
	Size          uint     `json:"size"`
}

func (f PuzzleFilter) Valid() error {
	if f.Themes == nil {
		return errors.New("themes cannot be nil, empty array should be used instead")
	}

	if f.MinRating < 0 {
		return errors.New("minimum rating must at least be 0")
	}

	if f.MaxRating < f.MinRating {
		return errors.New("maximum rating cannot be less than minimum rating")
	}

	if f.MinPopularity > 100 {
		return errors.New("minimum popularity can at most be the maximum possible popularity, which is 100")
	}

	if f.Size < 1 || f.Size > 1000 {
		return errors.New("size must be between 1 and 1000")
	}

	return nil
}

// RatingInterval encode the requested rating interval as a string.
func (f PuzzleFilter) RatingInterval() string {
	if f.MaxRating == 0 {
		return "*"
	}

	return fmt.Sprintf("%d - %d", f.MinRating, f.MaxRating)
}

func (f PuzzleFilter) String() string {
	return fmt.Sprintf(
		"PuzzleFilter(themes=%v, ratingInterval=%s, minPopularity=%d, size=%d)",
		f.Themes,
		f.RatingInterval(),
		f.MinPopularity,
		f.Size,
	)
}
