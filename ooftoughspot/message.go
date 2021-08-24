package ooftoughspot

import (
	"math/rand"
	"time"

	"github.com/FredBainbridge/PokerAPI/mavens/ringgame"
)

func SendOof() string {
	var taunts = [3]string{"Oof.", "Tough spot!", "Oof, tough spot!"}
	rand.Seed(time.Now().UnixNano())
	oof := taunts[rand.Intn(3)]
	ringgame.AdminMessage("Cash Game 10/20 No Limit ($1000 min)", oof)
	return oof
}
