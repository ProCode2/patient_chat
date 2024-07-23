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
