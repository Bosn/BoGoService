package main

import (
	"net/http"

	"github.com/bosn/BoGoService/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
