package store

import (
	"log"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	types "github.com/patient_chat/patient_chat_server/internal/data"
)

type DbStore struct {
	db *sqlx.DB
}

func NewStore() *DbStore {
	db, err := sqlx.Connect("sqlite3", "_patient.db")

	if err != nil {
		log.Fatal("Can not connect to db", err)
	}

	return &DbStore{db: db}
}

type GetSession struct {
	Sid    string `db:"session_id"`
	Uid    string `db:"user_id"`
	expiry string `db:"expired_at"`
}

func (d *DbStore) GetUserSession(uid string) (*types.Session, error) {
	s := `SELECT * FROM sessions WHERE user_id = ?`
	var ses GetSession
	res := d.db.QueryRow(s, uid)

	err := res.Scan(&ses)

	if err != nil {
		log.Println("Can not get sessinos of user")
		return nil, err
	}

	t, err := time.Parse(time.RFC1123, ses.expiry)
	if err != nil {
		log.Println("Can not get sessinos of user", err)
		return nil, err
	}

	session := &types.Session{
		UserID:    ses.Uid,
		SessionID: ses.Sid,
		ExpiresAt: t,
	}

	return session, nil

}

func (d *DbStore) CreateUserSession(uid, sid string, expiry time.Time) error {
	sSql := `INSERT INTO sessions (session_id, user_id, expires_at) VALUES (?, ?, ?)`
	_, err := d.db.Exec(sSql, sid, sid, expiry)
	if err != nil {
		log.Println("Can not create new session: ", err)
		return err
	}
	return nil
}

func (d *DbStore) DeleteUserSession(sid string) error {
	sSql := `DELETE FROM sessions WHERE sid = ?`
	_, err := d.db.Exec(sSql, sid)
	if err != nil {
		log.Println("Can not delete session: ", err)
		return err
	}
	return nil
}

func (d *DbStore) CreateNewUser(id, name, phone string, role types.Role) error {
	sSql := `INSERT INTO users (id, name, phone, role) VALUES (?, ?, ?, ?)`
	_, err := d.db.Exec(sSql, id, name, phone, role)
	if err != nil {
		log.Println("Can not create user: ", err)
		return err
	}
	return nil
}

func (d *DbStore) GetUserByPhone(phone string) (*types.User, error) {
	s := `SELECT id, name, phone, role FROM users WHERE phone = ?`
	var u types.User
	err := d.db.Get(&u, s, phone)
	if err != nil {
		log.Println("Can not get User by phone", err)
		return nil, err
	}

	return &u, nil
}

func (d *DbStore) CreateNewPatient(id, uid, did string, mhs []string) error {
	s := `INSERT INTO patients (id, user_id, doc_id, medical_history) VALUES (?, ?, ?, ?)`
	_, err := d.db.Exec(s, id, uid, did, strings.Join(mhs, ","))
	if err != nil {
		log.Println("Can not cretae patient: ", err)
		return err
	}
	return nil
}

func (d *DbStore) CreateNewDoc(id, uid, q, h string) error {
	s := `INSERT INTO doctors (id, user_id, qualification, hospital) VALUES (?, ?, ?, ?)`
	_, err := d.db.Exec(s, id, uid, q, h)
	if err != nil {
		log.Println("Can not cretae doctor: ", err)
		return err
	}
	return nil
}
