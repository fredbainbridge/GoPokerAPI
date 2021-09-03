package main

import (
	"net/http"

	"github.com/FredBainbridge/PokerAPI/controllers"
)

func main() {
	//data.Players()
	controllers.RegisterControllers()
	http.ListenAndServe(":8080", nil)
}
