package store

import (
	"time"

	types "github.com/patient_chat/patient_chat_server/internal/data"
)

type DbStorer interface {
	// session
	GetUserSessionBySessionID(sid string) (*types.Session, error)
	GetUserSession(uid string) (*types.Session, error)
	CreateUserSession(uid, sid string, expiry time.Time) error
	DeleteUserSession(sid string) error
	// user
	CreateNewUser(id, name, phone string, role types.Role) error
	GetUserByID(uid string) (*types.User, error)
	GetUserByPhone(phone string) (*types.User, error)
	// patient
	CreateNewPatient(id, uid, did string, mhs []string) error
	GetPatientByUserID(uid string) (*types.Patient, error)
	// doc
	CreateNewDoc(id, uid, q, h string) error
	GetDocByUserID(uid string) (*types.Doctor, error)
	GetAllDoctorUsers() ([]types.User, error)
	// chats
	GetChats(uid string) ([]types.Chat, error)
	AddChat(id, pid, did, tid, q, r string) error
	GetChatsByThreadID(tid string) ([]types.Chat, error)
}
