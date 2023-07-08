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

func (u *User) Deactivate() error {
	u.Active = false

	return nil
}
