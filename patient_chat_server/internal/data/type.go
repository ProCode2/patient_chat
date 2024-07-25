package types

import "time"

type Role string

const (
	DoctorType  Role = "doctor"
	PatientType Role = "patient"
)

type User struct {
	ID    string `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Phone string `json:"phone" db:"phone"`
	Role  Role   `json:"role" db:"role"`
}

type Session struct {
	UserID    string    `json:"userId"`
	SessionID string    `json:"sessionId"`
	ExpiresAt time.Time `json:"expiresAt"`
}

type Doctor struct {
	ID            string `json:"id" db:"id"`
	UserID        string `json:"userId" db:"user_id"`
	Qualification string `json:"qualification" db:"qualification"`
	Hospital      string `json:"hospital" db:"hospital"`
}

type Patient struct {
	ID             string   `json:"id" db:"id"`
	UserID         string   `json:"userId" db:"user_id"`
	DocID          string   `json:"docId" db:"doc_id"`
	MedicalHistory []string `json:"medicalHistory" db:"medical_history"`
}

type PatientUser struct {
	User    *User    `json:"user"`
	Patient *Patient `json:"patient"`
}

type DoctorUser struct {
	User   *User   `json:"user"`
	Doctor *Doctor `json:"doctor"`
}

type Chat struct {
	ID        string `json:"id" db:"ID"`
	PatientID string `json:"patientId" db:"patient_id"`
	DoctorID  string `json:"doctorId" db:"doc_id"`
	ThreadID  string `json:"threadId" db:"thread_id"`
	Query     string `json:"query" db:"query"`
	Response  string `json:"response" db:"response"`
	Time      string `json:"time" db:"time"`
}
