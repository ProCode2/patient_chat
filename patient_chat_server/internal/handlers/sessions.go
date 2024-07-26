package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/render"
	"github.com/patient_chat/patient_chat_server/internal/models"
)

func GetDoctorsHandler(w http.ResponseWriter, r *http.Request) {
	ds, err := models.GetDoctors()

	if err != nil {
		log.Println("can not get all doctors: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Respond(w, r, ds)

}

type SignUpRequestBody struct {
	PatientName string `json:"patientName"`
	DoctorID    string `json:"doctorId"`
	PhoneNumber string `json:"phoneNumber"`
	Otp         string `json:"otp"`
}

func (s *SignUpRequestBody) Bind(r *http.Request) error {
	if s.PatientName == "" || s.DoctorID == "" {
		return errors.New("Patient Name and Doctor Link is required")
	}

	if s.PhoneNumber == "" || len(s.PhoneNumber) != 10 {
		return errors.New("Invalid Phone number")
	}

	if s.Otp == "" || len(s.Otp) != 6 {
		return errors.New("Invalid OTP")
	}

	return nil
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var b SignUpRequestBody
	err := render.Bind(r, &b)
	if err != nil {
		log.Println("Binding err: ", err)
		w.WriteHeader(400)
		render.Respond(w, r, map[string]string{"message": err.Error()})
		return
	}

	// TODO: check if otp is correct

	p, err := models.SignUpPatient(b.PatientName, b.DoctorID, b.PhoneNumber)
	if err != nil {
		w.WriteHeader(500)
		render.Respond(w, r, map[string]string{"message": err.Error()})
		return
	}

	// create a new session
	s, err := models.CreateSessionForUser(p.UserID)
	if err != nil {
		w.WriteHeader(500)
		render.Respond(w, r, map[string]string{"message": err.Error()})
	}

	render.Respond(w, r, s)
}

type LoginRequestBody struct {
	Otp   string `json:"otp"`
	Phone string `json:"phone"`
}

func (l *LoginRequestBody) Bind(r *http.Request) error {
	if l.Otp == "" || l.Phone == "" || len(l.Phone) != 10 || len(l.Otp) != 6 {
		return errors.New("Phone and OTP is required")
	}
	return nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var b LoginRequestBody
	err := render.Bind(r, &b)
	if err != nil {
		http.Error(w, "Can not parse request data properly", http.StatusBadRequest)
		return
	}

	// TODO: check if otp is correct

	p, err := models.LoginPatient(b.Phone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// create new session for patient user
	s, err := models.CreateSessionForUser(p.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Respond(w, r, s)
}

// type LogOutRequestBody struct {
// 	SessionID string `json:"sessionId"`
// }

// func (l *LogOutRequestBody) Bind(r *http.Request) error {
// 	if len(l.SessionID) == 0 {
// 		return errors.New("Session Id is required to logout")
// 	}
// 	return nil
// }

func LogOutHandler(w http.ResponseWriter, r *http.Request) {
	p := r.Context().Value("patient")
	s := r.Context().Value("session").(string)
	if s == "" {
		http.Error(w, "Session data not found", http.StatusBadRequest)
		return
	}
	if p == nil {
		http.Error(w, "Patient data not found", http.StatusBadRequest)
		return
	}

	err := models.DeleteSessionForUser(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Respond(w, r, map[string]string{"message": "You have been succesfully logged out."})
}
