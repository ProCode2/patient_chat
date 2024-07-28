package routes

import (
	"net/http"
	"strings"

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
	r.Get("/docs", handlers.GetDoctorsHandler)
	r.Route("/", func(r chi.Router) {
		r.Use(middlewares.Authenticate)

		r.Delete("/logout", handlers.LogOutHandler)
		r.Route("/patient", func(r chi.Router) {
			r.Get("/", handlers.GetPatient)
			r.Put("/", handlers.UpdatePatientDataHandler)
			r.Get("/doc", handlers.GetPatientDoc)
			r.Get("/chats", handlers.GetChatsHandler)
			r.Get("/chats/{chatID}", handlers.GetChatsByThreadIDHandler)
			r.Post("/chats", handlers.AddChatHandler)
		})
	})
	return r
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
