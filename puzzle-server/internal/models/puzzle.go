package models

import (
	"fmt"
	"strings"
	"time"
)

const (
	separator = ","
)

// Puzzle represents a chess problem and solution along with associated metadata.
type Puzzle struct {
	ID              string    `json:"id"`
	ExternalID      string    `json:"externalId"`
	FEN             string    `json:"fen"`
	Moves           []string  `json:"moves"`
	Rating          int       `json:"rating"`
	RatingDeviation int       `json:"ratingDeviation"`
	Popularity      int       `json:"popularity"`
	Themes          []string  `json:"themes"`
	GameURL         string    `json:"gameUrl"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

// EncodeMoves converts moves into a string.
func (p Puzzle) EncodeMoves() string {
	return strings.Join(p.Moves, separator)
}

// EncodeThemes converts movves into a string.
func (p Puzzle) EncodeThemes() string {
	themes := make([]string, len(p.Themes))
	for i, theme := range p.Themes {
		themes[i] = encodeTheme(theme)
	}
	return strings.Join(themes, separator)
}

func (p Puzzle) String() string {
	return fmt.Sprintf(
		"Puzzle(id=%s, externalId=%s, rating=%d, popularity=%d)",
		p.ID,
		p.ExternalID,
		p.Rating,
		p.Popularity,
	)
}

func encodeTheme(theme string) string {
	return fmt.Sprintf("[%s]", theme)
}
