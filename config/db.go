package config

import (
	"gitlab.com/zikazama/golang-final-project/helpers"
	"gitlab.com/zikazama/golang-final-project/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB() {
	var err error
	// DB, err = gorm.Open("mysql", "root:@/egommerce?charset=utf8&parseTime=True&loc=Local")
	DB, err = gorm.Open("mysql", "sql6443357:lztBcCwQGV@tcp(sql6.freesqldatabase.com)/sql6443357?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Gagal konek ke db")
	}
	// defer DB.Close()

	Migrate()
}

func Migrate() {

	// migrasi
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Barang{})
	DB.AutoMigrate(&models.Kategori{})
	DB.AutoMigrate(&models.Pengirim{})
	DB.AutoMigrate(&models.Voucher{})
	DB.AutoMigrate(&models.Pembelian{})
	DB.AutoMigrate(&models.Detail_pembelian{})

	data := models.User{}
	if DB.Where("email = ? AND role = ?", "fauzi190198@gmail.com", "admin").Find(&data).RecordNotFound() {
		seederUser()
		seederKategori()
		seederPengirim()
		seederVoucher()
		seederBarang()
	}
}

func seederUser() {
	hashPassword, _ := helpers.HashPassword("123456")
	data := models.User{
		Email:    "fauzi190198@gmail.com",
		Password: hashPassword,
		Username: "zikazama",
		Role:     "admin",
		Nama:     "Fauzi",
		Alamat:   "Bandung",
	}
	DB.Create(&data)
}

func seederKategori() {
	kategori := []models.Kategori{
		{
			Nama_kategori: "Buku",
		}, {
			Nama_kategori: "Elektronik",
		},
		{
			Nama_kategori: "Gaming",
		},
		{
			Nama_kategori: "Kamera",
		},
		{
			Nama_kategori: "Fashion",
		},
	}
	for _, row := range kategori {
		DB.Create(&row)
	}
}

func seederPengirim() {
	pengirim := []models.Pengirim{
		{
			Nama_pengirim: "JNE",
		}, {
			Nama_pengirim: "Si Cepat",
		},
		{
			Nama_pengirim: "Anter aja",
		},
	}
	for _, row := range pengirim {
		DB.Create(&row)
	}
}

func seederVoucher() {
	voucher := []models.Voucher{
		{
			Nama_voucher:      "Voucher Senin",
			Deskripsi_voucher: "Hanya untuk hari tertentu",
			Diskon:            0.2,
		},
		{
			Nama_voucher:      "Voucher Selasa",
			Deskripsi_voucher: "Hanya untuk hari tertentu",
			Diskon:            0.3,
		},
		{
			Nama_voucher:      "Voucher Rabu",
			Deskripsi_voucher: "Hanya untuk hari tertentu",
			Diskon:            0.4,
		},
	}
	for _, row := range voucher {
		DB.Create(&row)
	}
}

func seederBarang() {
	barang := []models.Barang{
		{
			Nama_barang:  "Apple",
			Harga_barang: "8500000",
			Stok_barang:  10,
			Kategori_ID:  2,
		},
		{
			Nama_barang:  "Samsung",
			Harga_barang: "1000000",
			Stok_barang:  12,
			Kategori_ID:  2,
		},
		{
			Nama_barang:  "Asus ROG",
			Harga_barang: "25000000",
			Stok_barang:  9,
			Kategori_ID:  2,
		},
	}
	for _, row := range barang {
		DB.Create(&row)
	}
}
