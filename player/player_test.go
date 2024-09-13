package player_test

import (
	"testing"

	"blackjack/player"

	"github.com/stretchr/testify/assert"
)

func TestPlayer_GetName_ReturnsNamePassedThrough(t *testing.T) {
	testCases := map[string]struct {
		name string
	}{
		"Set name to Jane": {
			"Jane",
		},
		"Set name to Donald": {
			"Donald",
		},
	}

	for name, data := range testCases {
		t.Run(name, func(t *testing.T) {
			p := player.New(data.name)

			assert.Equal(t, data.name, p.Name())
		})
	}
}

func TestPlayer_GetHand_IsInitialised(t *testing.T) {
	p := player.New("Roger")

	assert.NotNil(t, p.Hand())
}
