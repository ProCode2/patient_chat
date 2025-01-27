package store

import (
	"log"
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

// get all users who are doctors
func (d *DbStore) GetAllDoctorUsers() ([]types.User, error) {
	s := `SELECT id, name FROM users WHERE role = "doctor"`

	var ds []types.User
	err := d.db.Select(&ds, s)
	if err != nil {
		return nil, err
	}

	return ds, nil
}

type GetSession struct {
	Sid    string `db:"session_id"`
	Uid    string `db:"user_id"`
	Expiry string `db:"expires_at"`
}

func (d *DbStore) GetUserSession(uid string) (*types.Session, error) {
	s := `SELECT * FROM sessions WHERE user_id = ?`
	var ses GetSession
	err := d.db.Get(&ses, s, uid)

	if err != nil {
		log.Println("Can not get sessions of user", err)
		return nil, err
	}

	t, err := time.Parse(time.RFC3339, ses.Expiry)
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

func (d *DbStore) GetUserSessionBySessionID(sid string) (*types.Session, error) {
	s := `SELECT * FROM sessions WHERE session_id = ?`
	var ses GetSession
	err := d.db.Get(&ses, s, sid)

	if err != nil {
		log.Println("Can not get sessinos of user", err)
		return nil, err
	}

	t, err := time.Parse(time.RFC3339, ses.Expiry)
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
	_, err := d.db.Exec(sSql, sid, uid, expiry.Format(time.RFC3339))
	if err != nil {
		log.Println("Can not create new session: ", err)
		return err
	}
	return nil
}

func (d *DbStore) DeleteUserSession(sid string) error {
	sSql := `DELETE FROM sessions WHERE session_id = ?`
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

func (d *DbStore) CreateNewPatient(id, uid, did, mhs string) error {
	s := `INSERT INTO patients (id, user_id, doc_id, medical_history) VALUES (?, ?, ?, ?)`
	_, err := d.db.Exec(s, id, uid, did, mhs)
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

func (d *DbStore) GetUserByID(uid string) (*types.User, error) {
	s := `SELECT id, name, phone, role FROM users WHERE id = ?`
	var u types.User
	err := d.db.Get(&u, s, uid)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (d *DbStore) GetPatientByUserID(uid string) (*types.Patient, error) {
	s := `SELECT id, user_id, doc_id, medical_history FROM patients WHERE user_id = ?`
	var p types.Patient

	err := d.db.Get(&p, s, uid)

	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (d *DbStore) GetDocByUserID(did string) (*types.Doctor, error) {
	s := `SELECT id, user_id, qualification, hospital FROM doctors WHERE user_id = ?`
	var dc types.Doctor
	err := d.db.Get(&dc, s, did)

	if err != nil {
		return nil, err
	}

	return &dc, nil
}

// update user whos role is patient
func (d *DbStore) UpdatePatientUser(u *types.User) (bool, error) {
	s := `UPDATE users SET name = ?, phone = ? WHERE id = ?`
	_, err := d.db.Exec(s, u.Name, u.Phone, u.ID)
	if err != nil {
		return false, nil
	}

	return true, nil
}

// update patient data in patient table
func (d *DbStore) UpdatePatientData(u *types.Patient) (bool, error) {
	log.Println(u.MedicalHistory)
	s := `UPDATE patients SET doc_id = ?, medical_history = ? WHERE user_id = ?`
	_, err := d.db.Exec(s, u.DocID, u.MedicalHistory, u.UserID)
	if err != nil {
		return false, nil
	}

	return true, nil
}
