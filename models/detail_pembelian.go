package models

import "gorm.io/gorm"

type Detail_pembelian struct {
	gorm.Model
	Pembelian_ID int
	Barang_ID    int
	Barang       Barang `gorm:"Foreignkey:Barang_ID;association_foreignkey:ID;"`
	Kuantitas    int
	Subtotal     int
}
