package routes

import (
	"gitlab.com/zikazama/golang-final-project/config"
	"gitlab.com/zikazama/golang-final-project/models"

	"github.com/gin-gonic/gin"
)

func MonitorTransaksi(c *gin.Context) {
	id := c.Param("id")
	data := []models.Pembelian{}

	if id != "" {
		data := models.Pembelian{}
		if config.DB.Preload("Detail_pembelian.Barang").Preload("User").Preload("Pengirim").Preload("Voucher").Where("id = ?", id).Find(&data).RecordNotFound() {
			c.JSON(404, gin.H{
				"message": "Transaksi tidak ditemukan",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Transaksi ditemukan",
			"data":    data,
		})
		return
	} else {
		config.DB.Preload("Detail_pembelian.Barang").Preload("User").Preload("Pengirim").Preload("Voucher").Find(&data)
		c.JSON(200, gin.H{
			"message": "Transaksi ditemukan",
			"data":    data,
		})
		return
	}
}
