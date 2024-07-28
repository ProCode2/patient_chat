package routes

import (
	"net/http"
	// "os"
	"path/filepath"
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
	// Create a route along /public that will serve contents from
	// the ./dist/ folder.
	// workDir, _ := os.Getwd()
	// filesDir := http.Dir(filepath.Join(workDir, "dist"))
	// FileServer(r, "/", filesDir)
	// Serve static files
	// fs := http.FileServer(http.Dir("./dist"))
	// r.Handle("/", fs)
	// r.Handle("/chat", fs)
	// r.Handle("/settings", fs)
	r.Route("/api", func(r chi.Router) {
		r.Post("/signup", handlers.SignUpHandler)
		r.Post("/login", handlers.LoginHandler)
		r.Get("/docs", handlers.GetDoctorsHandler)
		r.Group(func(auth chi.Router) {
			auth.Use(middlewares.Authenticate)

			auth.Delete("/logout", handlers.LogOutHandler)
			auth.Route("/patient", func(p chi.Router) {
				p.Get("/", handlers.GetPatient)
				p.Put("/", handlers.UpdatePatientDataHandler)
				p.Get("/doc", handlers.GetPatientDoc)
				p.Get("/chats", handlers.GetChatsHandler)
				p.Get("/chats/{chatID}", handlers.GetChatsByThreadIDHandler)
				p.Post("/chats", handlers.AddChatHandler)
			})
		})
	})

	// Serve static files
	staticHandler := http.FileServer(http.Dir("./dist"))
	r.Handle("/assets/*", staticHandler)

	// Catch-all route to serve index.html for React Router
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join("dist", "index.html"))
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
