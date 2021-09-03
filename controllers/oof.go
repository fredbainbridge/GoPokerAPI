package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/FredBainbridge/PokerAPI/data"
	"github.com/FredBainbridge/PokerAPI/mavens/ringgame"
	"github.com/FredBainbridge/PokerAPI/ooftoughspot"
)

type oofController struct {
	oofType *string
}

func (oc oofController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := r.FormValue("text")
	if len(params) == 0 {
		return
	}
	p := strings.Replace(params, "<", "", -1)
	p = strings.Replace(p, ">", "", -1)
	p = strings.Split(p, "|")[0]
	p = strings.Replace(p, "@", "", -1)
	user := data.Users(p)[0]
	seatedPlayers, err := ringgame.Playing()
	if err != nil {
		log.Fatalf("Cannot get seated players. Error: %q", err)
	}
	for i := range seatedPlayers {
		if user.UserName == seatedPlayers[i].Name {
			ooftoughspot.SendOof(seatedPlayers[i].TableName)
		}
	}
}
func NewOofController() *oofController {
	return &oofController{
		oofType: nil,
	}
}
