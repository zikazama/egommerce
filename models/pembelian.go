package models

import "gorm.io/gorm"

type Pembelian struct {
	gorm.Model
	User_ID          int
	User             User `gorm:"Foreignkey:User_ID;association_foreignkey:ID;"`
	Pengirim_ID      int
	Pengirim         Pengirim `gorm:"Foreignkey:Pengirim_ID;association_foreignkey:ID;"`
	Voucher_ID       int
	Voucher          Voucher `gorm:"Foreignkey:Voucher_ID;association_foreignkey:ID;"`
	Detail_pembelian []Detail_pembelian
	Total            int
}
