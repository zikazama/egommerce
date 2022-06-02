package models

import "gorm.io/gorm"

type Barang struct {
	gorm.Model
	Kategori_ID  int
	Kategori     Kategori `gorm:"Foreignkey:Kategori_ID;association_foreignkey:ID;"`
	Nama_barang  string
	Harga_barang string
	Stok_barang  int
}
