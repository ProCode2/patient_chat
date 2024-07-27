package handlers

import (
	"net/http"

	"github.com/go-chi/render"
	types "github.com/patient_chat/patient_chat_server/internal/data"
	"github.com/patient_chat/patient_chat_server/internal/models"
)

func GetPatient(w http.ResponseWriter, r *http.Request) {
	p := r.Context().Value("patient")

	p, ok := p.(*types.PatientUser)
	if !ok {
		http.Error(w, "Patient not found", http.StatusBadRequest)
		return
	}

	render.Respond(w, r, p)
}

func GetPatientDoc(w http.ResponseWriter, r *http.Request) {
	p := r.Context().Value("patient")

	up, ok := p.(*types.PatientUser)
	if !ok {
		http.Error(w, "Patient not found", http.StatusBadRequest)
		return
	}

	d, err := models.GetDocByUserID(up.Patient.DocID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Respond(w, r, d)
}

type UpdatePatientBody struct {
	DocID          string `json:"docId"`
	Name           string `json:"name"`
	MedicalHistory string `json:medicalHistory`
}

func (u *UpdatePatientBody) Bind(r *http.Request) error {
	return nil
}

func UpdatePatientDataHandler(w http.ResponseWriter, r *http.Request) {
	var u UpdatePatientBody

	err := render.Bind(r, &u)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p, ok := r.Context().Value("patient").(*types.PatientUser)

	if !ok {
		http.Error(w, "Patient not found", http.StatusBadRequest)
		return
	}

	d, err := models.UpdatePatient(p, u.DocID, u.Name, u.MedicalHistory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Respond(w, r, d)

}
