package models

import "gorm.io/gorm"

type Pengirim struct {
	gorm.Model
	Nama_pengirim string
}
