package users

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Fullname string
	Username string
	Active   bool
}
