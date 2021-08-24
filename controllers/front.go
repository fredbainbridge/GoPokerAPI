package controllers

import "net/http"

func RegisterControllers() {
	oc := NewOofController()
	http.Handle("/oof", *oc)
	http.Handle("/oof/", *oc)
}