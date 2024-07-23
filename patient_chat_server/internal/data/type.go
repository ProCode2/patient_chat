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
	ID            string `json:"id"`
	UserID        string `json:"userId"`
	Qualification string `json:"qualification"`
	Hospital      string `json:"hospital"`
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
