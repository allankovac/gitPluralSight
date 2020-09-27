package main

import (
	"net/http"

	"github.com/allankovac/gitPluralSight/controllers"
)

func main() {
	controllers.RegisterControllers()

	http.ListenAndServe(":3000", nil)
}
