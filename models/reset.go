package models

import uuid "github.com/satori/go.uuid"

type PasswordReset struct {
	Id    uuid.UUID
	Email string
	Token string
}
