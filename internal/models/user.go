package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
    ID       uint   `json:"id" gorm:"primaryKey"`
    Username string `json:"username" gorm:"unique;not null"`
    Password string `json:"password" gorm:"not null"`
    Role     string `json:"role" gorm:"default:'warga'"`
}

// Fungsi untuk meng-hash password
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
