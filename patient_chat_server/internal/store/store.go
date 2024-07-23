package store

import (
	"time"

	types "github.com/patient_chat/patient_chat_server/internal/data"
)

type DbStorer interface {
	GetUserSessionBySessionID(sid string) (*types.Session, error)
	GetUserSession(uid string) (*types.Session, error)
	CreateUserSession(uid, sid string, expiry time.Time) error
	DeleteUserSession(sid string) error
	CreateNewUser(id, name, phone string, role types.Role) error
	GetUserByID(uid string) (*types.User, error)
	GetUserByPhone(phone string) (*types.User, error)
	CreateNewPatient(id, uid, did string, mhs []string) error
	GetPatientByUserID(uid string) (*types.Patient, error)
	CreateNewDoc(id, uid, q, h string) error
	GetDocByUserID(uid string) (*types.Doctor, error)
}
