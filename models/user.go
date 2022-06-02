package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
	Role     string
	Nama     string
	Alamat   string
}
