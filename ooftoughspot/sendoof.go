package ooftoughspot

import (
	"math/rand"
	"time"

	"github.com/FredBainbridge/PokerAPI/mavens/ringgame"
)

func SendOof(tablename string) string {
	var taunts = [4]string{"OOF.", "Oof.", "Tough spot!", "Oof, tough spot!"}
	rand.Seed(time.Now().UnixNano())
	oof := taunts[rand.Intn(4)]
	ringgame.AdminMessage(tablename, oof)
	return oof
}
