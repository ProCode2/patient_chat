package models

import "github.com/google/uuid"

func GenID() string {
	return uuid.NewString()
}
