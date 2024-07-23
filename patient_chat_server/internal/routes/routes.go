package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/patient_chat/patient_chat_server/internal/handlers"
	"github.com/patient_chat/patient_chat_server/internal/middlewares"
)

func LoadRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Post("/signup", handlers.SignUpHandler)
	r.Post("/login", handlers.LoginHandler)
	r.Route("/", func(r chi.Router) {
		r.Use(middlewares.Authenticate)

		r.Delete("/logout", handlers.LogOutHandler)
	})
	return r
}
