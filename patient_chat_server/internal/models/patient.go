package models

import (
	"errors"
	types "github.com/patient_chat/patient_chat_server/internal/data"
	"log"
)

func GetDocByUserID(uid string) (*types.DoctorUser, error) {
	d, err := db.GetDocByUserID(uid)
	if err != nil {
		log.Println("Can not get Doctor by ID", err)
		return nil, errors.New("Something went wrong while getting doctor by id")
	}

	u, err := db.GetUserByID(d.UserID)

	if err != nil {
		log.Println("Can not get User by ID", err)
		return nil, errors.New("Something went wrong while getting user by ID")
	}
	ud := &types.DoctorUser{
		User:   u,
		Doctor: d,
	}

	return ud, nil
}

func UpdatePatient(p *types.PatientUser, did, name, mhs string) (bool, error) {
	if len(name) == 0 || len(did) == 0 {
		return false, errors.New("Insufficient data.")
	}

	p.User.Name = name
	p.Patient.DocID = did
	if p.Patient.MedicalHistory != "" {
		p.Patient.MedicalHistory += ("," + mhs)
	}
	p.Patient.MedicalHistory += mhs
	_, err := db.UpdatePatientUser(p.User)
	if err != nil {
		log.Println("can not update patient user: ", err)
		return false, errors.New("Something went wrong while updaing user")
	}

	_, err = db.UpdatePatientData(p.Patient)

	if err != nil {
		log.Println("can not update patient data: ", err)
		return false, errors.New("Something went wrong while updaing patient data")
	}

	return true, nil

}
