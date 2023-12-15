package models

import (
	"github.com/google/uuid"
)

/*
RegisterUser : Request
Login : Request
*/
type AuthUser struct {
	Email    string
	Password string
}

/*
Login : Response
*/
type UUID struct {
	Uuid uuid.UUID
}
