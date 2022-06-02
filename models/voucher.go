package models

import "gorm.io/gorm"

type Voucher struct {
	gorm.Model
	Nama_voucher      string
	Deskripsi_voucher string
	Diskon            float64
}
