package controllers

import (
	"net/http"

	"github.com/FredBainbridge/PokerAPI/ooftoughspot"
)

type oofController struct {
	oofType *string
}

func (oc oofController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(ooftoughspot.SendOof()))
}
func NewOofController() *oofController {
	return &oofController{
		oofType: nil,
	}
}
