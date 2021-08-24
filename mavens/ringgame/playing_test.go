package ringgame_test

import (
	"testing"

	"github.com/FredBainbridge/PokerAPI/mavens/ringgame"
)

func TestPlaying(t *testing.T) {
	players, err := ringgame.Playing()
	if(err != nil) {
		t.Errorf("Failed with error: %q", err)
	}
	for i := range players {
		if(len(players[i].Name) == 0 ) {
			t.Errorf("Empty player name.")
		} 
	}
}