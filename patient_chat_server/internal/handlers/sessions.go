package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/render"
	"github.com/patient_chat/patient_chat_server/internal/models"
)

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
		render.Respond(w, r, map[string]string{"message": "Bad Data"})
		return
	}

	// TODO: check if otp is correct

	p, err := models.SignUpPatient(b.PatientName, b.DoctorID, b.PhoneNumber)
	if err != nil {
		w.WriteHeader(500)
		render.Respond(w, r, map[string]string{"message": "Something went wrong."})
		return
	}

	// create a new session
	s, err := models.CreateSessionForUser(p.UserID)
	if err != nil {
		w.WriteHeader(500)
		render.Respond(w, r, map[string]string{"message": "Something went wrong."})
	}

	render.Respond(w, r, s)
}
