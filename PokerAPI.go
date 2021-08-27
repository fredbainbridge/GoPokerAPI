package main

import (
	"net/http"

	"github.com/FredBainbridge/PokerAPI/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":8080", nil)
}
