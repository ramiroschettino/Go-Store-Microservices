package domain

import (
    "errors"
    "golang.org/x/crypto/bcrypt"
)

type User struct {
    ID       string `json:"id"`
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=8"`
}

func (u *User) HashPassword() error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
    if err != nil {
        return errors.New("failed to hash password")
    }
    u.Password = string(hashedPassword)
    return nil
}

func (u *User) CheckPassword(password string) error {
    return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}