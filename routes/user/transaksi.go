package routes

import (
	"strconv"

	"gitlab.com/zikazama/golang-final-project/config"
	"gitlab.com/zikazama/golang-final-project/models"

	"github.com/gin-gonic/gin"
)

func CreateTransaksi(c *gin.Context) {
	barang := c.PostFormArray("barang_id")
	kuantitas := c.PostFormArray("kuantitas")
	var total int
	modelDetailPembelian := []models.Detail_pembelian{}
	for i, _ := range barang {
		modelBarang := models.Barang{}
		if config.DB.Preload("Kategori").Where("id = ?", barang[i]).Find(&modelBarang).RecordNotFound() {
			c.JSON(404, gin.H{
				"message": "Barang tidak ditemukan",
			})
			return
		}
		parse_barang_id, _ := strconv.Atoi(barang[i])
		parse_kuantitas, _ := strconv.Atoi(kuantitas[i])
		parse_subtotoal, _ := strconv.Atoi(modelBarang.Harga_barang)
		modelDetailPembelian = append(modelDetailPembelian, models.Detail_pembelian{
			Barang_ID: parse_barang_id,
			Kuantitas: parse_kuantitas,
			Subtotal:  parse_subtotoal,
		})
		total += parse_subtotoal
	}

	voucher_id := c.PostForm("voucher_id")
	if voucher_id == "" {
		parse_pengirim_id, _ := strconv.Atoi(c.PostForm("pengirim_id"))
		if config.DB.Where("id = ?", parse_pengirim_id).Find(&models.Pengirim{}).RecordNotFound() {
			c.JSON(404, gin.H{
				"message": "Pengirim tidak ditemukan",
			})
			return
		}

		user_id, _ := c.Get("user_id")
		parse_user_id := user_id.(float64)
		parse_user_id_int := int(parse_user_id)
		modelPembelian := models.Pembelian{
			User_ID:     parse_user_id_int,
			Pengirim_ID: parse_pengirim_id,
			Total:       total,
		}
		config.DB.Create(&modelPembelian)

		for i, _ := range modelDetailPembelian {
			modelDetailPembelian[i].Pembelian_ID = int(modelPembelian.ID)
			config.DB.Create(&modelDetailPembelian[i])
		}

		c.JSON(201, gin.H{
			"message": "Berhasil melakukan transaksi",
			"data":    modelPembelian,
		})
	} else {
		modelVoucher := models.Voucher{}
		if config.DB.Where("id = ?", voucher_id).Find(&modelVoucher).RecordNotFound() {
			c.JSON(404, gin.H{
				"message": "Voucher tidak ditemukan",
			})
			return
		}

		// penerapan diskon
		total = total - int(modelVoucher.Diskon*float64(total))

		parse_pengirim_id, _ := strconv.Atoi(c.PostForm("pengirim_id"))
		if config.DB.Where("id = ?", parse_pengirim_id).Find(&models.Pengirim{}).RecordNotFound() {
			c.JSON(404, gin.H{
				"message": "Pengirim tidak ditemukan",
			})
			return
		}

		user_id, _ := c.Get("user_id")
		parse_user_id := user_id.(float64)
		parse_user_id_int := int(parse_user_id)
		parse_voucher_id, _ := strconv.Atoi(voucher_id)
		modelPembelian := models.Pembelian{
			User_ID:     parse_user_id_int,
			Pengirim_ID: parse_pengirim_id,
			Voucher_ID:  parse_voucher_id,
			Total:       total,
		}
		config.DB.Create(&modelPembelian)

		for i, _ := range modelDetailPembelian {
			modelDetailPembelian[i].Pembelian_ID = int(modelPembelian.ID)
			config.DB.Create(&modelDetailPembelian[i])
			modelBarang := models.Barang{}
			if config.DB.Preload("Kategori").Where("id = ?", modelDetailPembelian[i].Barang_ID).Find(&modelBarang).RecordNotFound() {
				c.JSON(404, gin.H{
					"message": "Barang tidak ditemukan",
				})
				return
			}
			// update stok barang
			dataUpdate := models.Barang{
				Nama_barang:  modelBarang.Nama_barang,
				Harga_barang: modelBarang.Harga_barang,
				Stok_barang:  modelBarang.Stok_barang - modelDetailPembelian[i].Kuantitas,
				Kategori_ID:  modelBarang.Kategori_ID,
			}
			config.DB.Model(&modelBarang).Updates(dataUpdate)
		}

		c.JSON(201, gin.H{
			"message": "Berhasil melakukan transaksi",
			"data":    modelPembelian,
		})
	}
}

func ReadTransaksi(c *gin.Context) {
	user_id, _ := c.Get("user_id")
	parse_user_id := user_id.(float64)
	parse_user_id_int := int(parse_user_id)
	id := c.Param("id")
	data := []models.Pembelian{}

	if id != "" {
		data := models.Pembelian{}
		if config.DB.Preload("Detail_pembelian.Barang").Preload("User").Preload("Pengirim").Preload("Voucher").Where("id = ? And User_ID = ?", id, parse_user_id_int).Find(&data).RecordNotFound() {
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
		config.DB.Preload("Detail_pembelian.Barang").Preload("User").Preload("Pengirim").Preload("Voucher").Where("User_ID = ?", parse_user_id_int).Find(&data)
		c.JSON(200, gin.H{
			"message": "Transaksi ditemukan",
			"data":    data,
		})
		return
	}
}
