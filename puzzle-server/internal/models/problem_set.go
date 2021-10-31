package models

import (
	"fmt"
	"time"
)

// ProblemSet represents a collection of chess puzzles
type ProblemSet struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Themes         []string  `json:"themes"`
	RatingInterval string    `json:"ratingInterval"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	PuzzleIDs      []string  `json:"puzzleIds"`
}

func (p ProblemSet) String() string {
	return fmt.Sprintf(
		"ProblemSet(id=%s, name=%s, ratingInterval=%s)",
		p.ID,
		p.Name,
		p.RatingInterval,
	)
}
