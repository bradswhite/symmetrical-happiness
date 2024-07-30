package types

import (
	"time"

	"github.com/google/uuid"
  "golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
  ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Token    string    `json:"token"`
}

type CreateUserRequest struct {
	Email     string `json:"email"`
  Username  string `json:"username"`
	Password  string `json:"password"`
}

type User struct {
	ID                uuid.UUID `json:"id"`
	Email             string    `json:"email"`
	Username          string    `json:"username"`
	EncryptedPassword string    `json:"-"`
	CreatedAt         time.Time `json:"createdAt"`
}

func (a *User) ValidPassword(pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(a.EncryptedPassword), []byte(pw)) == nil
}

func NewUser(email, username, password string) (*User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		Email:             email,
		Username:          username,
		EncryptedPassword: string(encpw),
		CreatedAt:         time.Now().UTC(),
	}, nil
}
