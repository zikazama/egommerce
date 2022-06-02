package models

import "gorm.io/gorm"

type Kategori struct {
	gorm.Model
	Nama_kategori string
}
