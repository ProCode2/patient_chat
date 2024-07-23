package middlewares

import (
	"context"
	"log"
	"net/http"

	types "github.com/patient_chat/patient_chat_server/internal/data"
	"github.com/patient_chat/patient_chat_server/internal/models"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a := r.Header.Get("Authentication")

		log.Println("Hello here", a)
		s, err := models.GetSessionBySessionID(a)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		u, err := models.GetUserFromID(s.UserID)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Println("uid", u.ID)
		p, err := models.GetPatientData(u.ID)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		up := &types.PatientUser{
			User:    u,
			Patient: p,
		}
		ctx := context.WithValue(r.Context(), "patient", up)
		ctx = context.WithValue(ctx, "session", s.SessionID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
