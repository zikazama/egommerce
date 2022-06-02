package main

import (
	"gitlab.com/zikazama/golang-final-project/config"
	"gitlab.com/zikazama/golang-final-project/middlewares"
	routes_admin "gitlab.com/zikazama/golang-final-project/routes/admin"
	routes_user "gitlab.com/zikazama/golang-final-project/routes/user"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	defer config.DB.Close()

	router := gin.Default()

	router.GET("/", index)

	// auth admin
	router.POST("/admin/login", routes_admin.Login)

	// auth user
	router.POST("/login", routes_user.Login)
	router.POST("/register", routes_user.Register)

	// route monitor transaksi
	router.GET("/admin/transaksi", middlewares.IsAuth("admin"), routes_admin.MonitorTransaksi)
	router.GET("/admin/transaksi/:id", middlewares.IsAuth("admin"), routes_admin.MonitorTransaksi)

	// route profile user
	router.PUT("/profile", middlewares.IsAuth("user"), routes_user.UpdateProfile)
	router.GET("/profile", middlewares.IsAuth("user"), routes_user.ReadProfile)

	// route transaksi
	router.POST("/transaksi", middlewares.IsAuth("user"), routes_user.CreateTransaksi)
	router.GET("/transaksi", middlewares.IsAuth("user"), routes_user.ReadTransaksi)
	router.GET("/transaksi/:id", middlewares.IsAuth("user"), routes_user.ReadTransaksi)

	// route barang
	router.POST("/barang", middlewares.IsAuth("admin"), routes_admin.StoreBarang)
	router.GET("/barang", middlewares.IsAuth("admin"), routes_admin.ReadBarang)
	router.GET("/barang/:id", middlewares.IsAuth("admin"), routes_admin.ReadBarang)
	router.PUT("/barang/:id", middlewares.IsAuth("admin"), routes_admin.UpdateBarang)
	router.DELETE("/barang/:id", middlewares.IsAuth("admin"), routes_admin.DeleteBarang)

	// route voucher
	router.POST("/voucher", middlewares.IsAuth("admin"), routes_admin.StoreVoucher)
	router.GET("/voucher", middlewares.IsAuth("admin"), routes_admin.ReadVoucher)
	router.GET("/voucher/:id", middlewares.IsAuth("admin"), routes_admin.ReadVoucher)
	router.PUT("/voucher/:id", middlewares.IsAuth("admin"), routes_admin.UpdateVoucher)
	router.DELETE("/voucher/:id", middlewares.IsAuth("admin"), routes_admin.DeleteVoucher)

	// route kategori
	router.POST("/kategori", middlewares.IsAuth("admin"), routes_admin.StoreKategori)
	router.GET("/kategori", middlewares.IsAuth("admin"), routes_admin.ReadKategori)
	router.GET("/kategori/:id", middlewares.IsAuth("admin"), routes_admin.ReadKategori)
	router.PUT("/kategori/:id", middlewares.IsAuth("admin"), routes_admin.UpdateKategori)
	router.DELETE("/kategori/:id", middlewares.IsAuth("admin"), routes_admin.DeleteKategori)

	// route pengirim
	router.POST("/pengirim", middlewares.IsAuth("admin"), routes_admin.StorePengirim)
	router.GET("/pengirim", middlewares.IsAuth("admin"), routes_admin.ReadPengirim)
	router.GET("/pengirim/:id", middlewares.IsAuth("admin"), routes_admin.ReadPengirim)
	router.PUT("/pengirim/:id", middlewares.IsAuth("admin"), routes_admin.UpdatePengirim)
	router.DELETE("/pengirim/:id", middlewares.IsAuth("admin"), routes_admin.DeletePengirim)

	// route user
	router.POST("/user", middlewares.IsAuth("admin"), routes_admin.StoreUser)
	router.GET("/user", middlewares.IsAuth("admin"), routes_admin.ReadUser)
	router.GET("/user/:id", middlewares.IsAuth("admin"), routes_admin.ReadUser)
	router.PUT("/user/:id", middlewares.IsAuth("admin"), routes_admin.UpdateUser)
	router.DELETE("/user/:id", middlewares.IsAuth("admin"), routes_admin.DeleteUser)

	router.Run()
}

func index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Halo ini API egommerce by Fauzi",
	})
}
