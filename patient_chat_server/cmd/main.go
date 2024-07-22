package main

import (
	"net/http"

	"github.com/patient_chat/patient_chat_server/internal/routes"
)

func main() {
	r := routes.LoadRoutes()
	http.ListenAndServe(":3333", r)
}
