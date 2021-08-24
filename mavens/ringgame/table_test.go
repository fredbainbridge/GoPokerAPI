package ringgame_test

import (
	"testing"

	"github.com/FredBainbridge/PokerAPI/mavens/ringgame"
)

func TestTablesExist(t *testing.T) {
	results, err := ringgame.Tables()
	if(err != nil) {
		t.Error(err)
	}
	if(len(results) == 0) {
		t.Error("No tables returned.")
	}
}