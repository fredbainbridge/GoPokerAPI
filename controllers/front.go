package controllers

import (
	"net/http"
)

//import "net/http"

func RegisterControllers() {
	oc := NewOofController()
	http.Handle("/oof", *oc)
	http.Handle("/oof/", *oc)
}
