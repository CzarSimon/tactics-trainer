package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzle_EncodeMoves(t *testing.T) {
	assert := assert.New(t)

	puzzle := Puzzle{
		ID: "some-id",
		Moves: []string{
			"e2e4",
			"e7e5",
			"g1f3",
		},
	}

	assert.Equal("e2e4,e7e5,g1f3", puzzle.EncodeMoves())

	puzzle.Moves = []string{}
	assert.Equal("", puzzle.EncodeMoves())
}

func TestPuzzle_EncodeThemes(t *testing.T) {
	assert := assert.New(t)

	puzzle := Puzzle{
		ID: "some-id",
		Themes: []string{
			"mateIn2",
			"short",
			"sacrifice",
		},
	}

	assert.Equal("[mateIn2],[short],[sacrifice]", puzzle.EncodeThemes())

	puzzle.Themes = []string{}
	assert.Equal("", puzzle.EncodeThemes())
}
