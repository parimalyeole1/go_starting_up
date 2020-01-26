package main

import (
	"net/http"

	"github.com/pluralsight/webservice/controllers"
)

func main() {
	controllers.RegisterController()
	http.ListenAndServe(":300", nil)
}
