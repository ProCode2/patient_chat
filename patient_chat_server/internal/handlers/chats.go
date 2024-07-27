package handlers

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	types "github.com/patient_chat/patient_chat_server/internal/data"
	"github.com/patient_chat/patient_chat_server/internal/models"
)

func GetChatsHandler(w http.ResponseWriter, r *http.Request) {
	p, ok := r.Context().Value("patient").(*types.PatientUser)
	if !ok {
		http.Error(w, "Patient not found.", http.StatusBadRequest)
		return
	}

	cs, err := models.GetChatsForUser(p.User.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.Respond(w, r, cs)
}

type AddChatRequestBody struct {
	Query    string `json:"query"`
	ThreadID string `json:"threadId"`
}

func (a *AddChatRequestBody) Bind(r *http.Request) error {
	if a.Query == "" {
		return errors.New("Query can not be empty")
	}
	return nil
}

func AddChatHandler(w http.ResponseWriter, r *http.Request) {
	var b AddChatRequestBody
	err := render.Bind(r, &b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p, ok := r.Context().Value("patient").(*types.PatientUser)
	if !ok {
		http.Error(w, "Patient not found.", http.StatusBadRequest)
		return
	}

	pid := p.User.ID
	did := p.Patient.DocID

	res, err := models.AddUserChat(pid, did, b.ThreadID, b.Query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.Respond(w, r, map[string]string{"threadId": res})
}

func GetChatsByThreadIDHandler(w http.ResponseWriter, r *http.Request) {
	tid := chi.URLParam(r, "chatID")

	cs, err := models.GetChatsByThreadID(tid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.Respond(w, r, cs)
}
