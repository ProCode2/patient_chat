package store

import (
	"time"

	types "github.com/patient_chat/patient_chat_server/internal/data"
)

type DbStorer interface {
	GetUserSession(uid string) (*types.Session, error)
	CreateUserSession(uid, sid string, expiry time.Time) error
	DeleteUserSession(sid string) error
	CreateNewUser(id, name, phone string, role types.Role) error
	GetUserByPhone(phone string) (*types.User, error)
	CreateNewPatient(id, uid, did string, mhs []string) error
	CreateNewDoc(id, uid, q, h string) error
}
