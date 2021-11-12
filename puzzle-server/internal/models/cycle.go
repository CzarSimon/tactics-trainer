package models

import (
	"fmt"
	"time"
)

// Cycle represents a round of going through the puzzles in a problem set
type Cycle struct {
	ID              string
	Number          int
	ProblemSetID    string
	CurrentPuzzleID string
	CompleatedAt    time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// Compleated returns true if the cycle has been marked as finished
func (c Cycle) Compleated() bool {
	return c.CompleatedAt.After(c.CreatedAt)
}

func (c Cycle) String() string {
	return fmt.Sprintf(
		"Cycle(id=%s, number=%d, problemSetID=%s currentPuzzleID=%s, createdAt=%v, compleated=%t)",
		c.ID,
		c.Number,
		c.ProblemSetID,
		c.CurrentPuzzleID,
		c.CreatedAt,
		c.Compleated(),
	)
}
