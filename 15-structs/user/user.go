package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

func New(name, email, password string) (*User, error) {
	if password == "" {
		return nil, errors.New("password must not be blank")
	}

	return &User{
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
	}, nil
}

func (u User) ShowDetails() {
	fmt.Printf("These are the user details: %v, %v, %v\n", u.Name, u.Email, u.Password)
}

func (u *User) UpdatePassword() {
	u.Password = "NewPassword"
}

// Inheritance
type Admin struct {
	GhToken string
	User
}

func NewAdmin(name, email, password, ghToken string) *Admin {
	return &Admin{
		GhToken: ghToken,
		User: User{
			Name:      name,
			Email:     email,
			Password:  password,
			CreatedAt: time.Now(),
		},
	}
}
