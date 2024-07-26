package models

import (
	"database/sql"
	"errors"
	"log"
	"time"

	types "github.com/patient_chat/patient_chat_server/internal/data"
	"github.com/patient_chat/patient_chat_server/internal/store"
)

var db store.DbStorer = store.NewStore()

func NewUser(name string, role types.Role, phone string) *types.User {
	_, err := db.GetUserByPhone(phone)

	if err != nil && err != sql.ErrNoRows {
		log.Println("Can not create user: %w", err)
		return nil
	}

	if err == sql.ErrNoRows {
		u := &types.User{
			ID:    GenID(),
			Name:  name,
			Role:  role,
			Phone: phone,
		}
		err = db.CreateNewUser(u.ID, u.Name, u.Phone, u.Role)
		if err != nil {
			log.Println("Can not create user: %w", err)
			return nil
		}
		return u
	}
	return nil
}

func NewSession(uid string) *types.Session {
	s := &types.Session{
		UserID:    uid,
		SessionID: GenID(),
		ExpiresAt: time.Now().Add(time.Hour * 24),
	}
	err := db.CreateUserSession(s.UserID, s.SessionID, s.ExpiresAt)
	if err != nil {
		log.Println("Can not create session: ", err)
		return nil
	}
	return s
}

func NewPatient(uid, did string) *types.Patient {
	p := &types.Patient{
		ID:             GenID(),
		UserID:         uid,
		DocID:          did,
		MedicalHistory: []string{},
	}

	err := db.CreateNewPatient(p.ID, p.UserID, p.DocID, p.MedicalHistory)

	if err != nil {
		log.Println("Can no create patient: %w", err)
		return nil
	}

	return p
}

func SignUpPatient(name, did, phone string) (*types.Patient, error) {
	u := NewUser(name, types.PatientType, phone)
	if u == nil {
		return nil, errors.New("Something went wrong while creating user.")
	}

	p := NewPatient(u.ID, did)
	if p == nil {
		return nil, errors.New("Something went wrong while creating patient.")
	}

	return p, nil
}

func LoginPatient(phone string) (*types.User, error) {
	u, err := db.GetUserByPhone(phone)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("No user with that phone number exists yet")
		} else {
			return nil, errors.New("Something went wrong while fetching user data")
		}
	}
	return u, nil

}

func CreateSessionForUser(uid string) (*types.Session, error) {
	prev, err := db.GetUserSession(uid)
	if err != nil && err != sql.ErrNoRows {
		return nil, errors.New("Something went wrong while fetching previous session.")
	}

	if err != sql.ErrNoRows {
		log.Println("SESSION", prev)
		err = db.DeleteUserSession(prev.SessionID)
		if err != nil {
			return nil, errors.New("Something went wrong while deleting previous session.")
		}
	}

	s := NewSession(uid)
	if s == nil {
		return nil, errors.New("Something went wrong while creating session.")
	}

	return s, nil
}

func DeleteSessionForUser(sid string) error {
	err := db.DeleteUserSession(sid)
	if err != nil {
		log.Println("Something went wrong when deleting session: ", err)
		return errors.New("Something went wrong when deleting session")
	}
	return nil
}

func GetSessionForUser(sid string) (*types.Session, error) {
	s, err := db.GetUserSession(sid)
	if err != nil {
		log.Println("Can not get user session", err)
		return nil, errors.New("Something went wrong when getting session")
	}

	return s, nil
}

func GetSessionBySessionID(sid string) (*types.Session, error) {
	s, err := db.GetUserSessionBySessionID(sid)
	if err != nil {
		log.Println("Can not get user session by ID", err)
		return nil, errors.New("Something went wrong when getting session")
	}

	return s, nil
}

func GetUserFromID(uid string) (*types.User, error) {
	u, err := db.GetUserByID(uid)
	if err != nil {
		log.Println("Can not get user with ID", err)
		return nil, errors.New("Something went wrong while fetching user with ID")
	}

	return u, nil
}

func GetPatientData(uid string) (*types.Patient, error) {
	p, err := db.GetPatientByUserID(uid)
	if err != nil {
		log.Println("Can not get patient with User ID", err)
		return nil, errors.New("Something went wrong while fetching patient with user ID")
	}
	return p, nil
}

func GetDoctors() ([]types.User, error) {
	ds, err := db.GetAllDoctorUsers()
	if err != nil {
		log.Println("Can not get all doctors", err)
		return nil, errors.New("Something went wrong while getting all doctors")
	}
	return ds, nil
}
