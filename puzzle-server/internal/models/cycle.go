package models

import (
	"fmt"
	"time"
)

// Cycle represents a round of going through the puzzles in a problem set
type Cycle struct {
	ID              string    `json:"id"`
	Number          int       `json:"number"`
	ProblemSetID    string    `json:"problemSetId"`
	CurrentPuzzleID string    `json:"currentPuzzleId"`
	CompletedAt     time.Time `json:"completedAt,omitempty"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

// Completed returns true if the cycle has been marked as finished
func (c Cycle) Completed() bool {
	return c.CompletedAt.After(c.CreatedAt)
}

func (c Cycle) String() string {
	return fmt.Sprintf(
		"Cycle(id=%s, number=%d, problemSetID=%s currentPuzzleID=%s, createdAt=%v, completed=%t)",
		c.ID,
		c.Number,
		c.ProblemSetID,
		c.CurrentPuzzleID,
		c.CreatedAt,
		c.Completed(),
	)
}
