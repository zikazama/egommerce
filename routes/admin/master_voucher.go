package routes

import (
	"strconv"

	"gitlab.com/zikazama/golang-final-project/config"
	"gitlab.com/zikazama/golang-final-project/models"

	"github.com/gin-gonic/gin"
)

func ReadVoucher(c *gin.Context) {
	id := c.Param("id")
	data := []models.Voucher{}

	if id != "" {
		data := models.Voucher{}
		if config.DB.Where("id = ?", id).Find(&data).RecordNotFound() {
			c.JSON(404, gin.H{
				"message": "Voucher tidak ditemukan",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Voucher ditemukan",
			"data":    data,
		})
		return
	} else {
		config.DB.Find(&data)
		c.JSON(200, gin.H{
			"message": "Voucher ditemukan",
			"data":    data,
		})
		return
	}
}

func StoreVoucher(c *gin.Context) {
	diskon, _ := strconv.ParseFloat(c.PostForm("diskon"), 64)
	data := models.Voucher{
		Nama_voucher:      c.PostForm("nama_voucher"),
		Deskripsi_voucher: c.PostForm("deskripsi_voucher"),
		Diskon:            diskon,
	}

	config.DB.Create(&data)

	c.JSON(201, gin.H{
		"message": "Berhasil menambahkan voucher",
		"data":    data,
	})
}

func UpdateVoucher(c *gin.Context) {
	id := c.Param("id")
	var data models.Voucher
	if config.DB.Where("id = ?", id).Find(&data).RecordNotFound() {
		c.JSON(404, gin.H{
			"message": "Voucher tidak ditemukan",
		})
		return
	}

	diskon, _ := strconv.ParseFloat(c.PostForm("diskon"), 64)
	dataUpdate := models.Voucher{
		Nama_voucher:      c.PostForm("nama_voucher"),
		Deskripsi_voucher: c.PostForm("deskripsi_voucher"),
		Diskon:            diskon,
	}

	config.DB.Model(&data).Updates(dataUpdate)

	c.JSON(200, gin.H{
		"message": "Berhasil memperbarui voucher",
		"data":    data,
	})
}

func DeleteVoucher(c *gin.Context) {
	id := c.Param("id")
	var data models.Voucher
	if config.DB.Where("id = ?", id).Find(&data).RecordNotFound() {
		c.JSON(404, gin.H{
			"message": "Voucher tidak ditemukan",
		})
		return
	}

	config.DB.Delete(&models.Voucher{}, id)

	c.JSON(200, gin.H{
		"message": "Berhasil menghapus voucher",
	})
}
