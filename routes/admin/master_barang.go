package routes

import (
	"strconv"

	"gitlab.com/zikazama/golang-final-project/config"
	"gitlab.com/zikazama/golang-final-project/models"

	"github.com/gin-gonic/gin"
)

func ReadBarang(c *gin.Context) {
	id := c.Param("id")
	data := []models.Barang{}

	if id != "" {
		data := models.Barang{}
		if config.DB.Preload("Kategori").Where("id = ?", id).Find(&data).RecordNotFound() {
			c.JSON(404, gin.H{
				"message": "Barang tidak ditemukan",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Barang ditemukan",
			"data":    data,
		})
		return
	} else {
		config.DB.Preload("Kategori").Find(&data)
		c.JSON(200, gin.H{
			"message": "Barang ditemukan",
			"data":    data,
		})
		return
	}
}

func StoreBarang(c *gin.Context) {
	kategori_id, _ := strconv.Atoi(c.PostForm("kategori_id"))
	stok_barang, _ := strconv.Atoi(c.PostForm("stok_barang"))
	data := models.Barang{
		Nama_barang:  c.PostForm("nama_barang"),
		Harga_barang: c.PostForm("harga_barang"),
		Stok_barang:  stok_barang,
		Kategori_ID:  kategori_id,
	}

	config.DB.Create(&data)

	c.JSON(201, gin.H{
		"message": "Berhasil menambahkan barang",
		"data":    data,
	})
}

func UpdateBarang(c *gin.Context) {
	id := c.Param("id")
	var data models.Barang
	if config.DB.Where("id = ?", id).Find(&data).RecordNotFound() {
		c.JSON(404, gin.H{
			"message": "Barang tidak ditemukan",
		})
		return
	}

	kategori_id, _ := strconv.Atoi(c.PostForm("kategori_id"))
	stok_barang, _ := strconv.Atoi(c.PostForm("stok_barang"))
	dataUpdate := models.Barang{
		Nama_barang:  c.PostForm("nama_barang"),
		Harga_barang: c.PostForm("harga_barang"),
		Stok_barang:  stok_barang,
		Kategori_ID:  kategori_id,
	}

	config.DB.Model(&data).Updates(dataUpdate)

	c.JSON(200, gin.H{
		"message": "Berhasil memperbarui barang",
		"data":    data,
	})
}

func DeleteBarang(c *gin.Context) {
	id := c.Param("id")
	var data models.Barang
	if config.DB.Where("id = ?", id).Find(&data).RecordNotFound() {
		c.JSON(404, gin.H{
			"message": "Barang tidak ditemukan",
		})
		return
	}

	config.DB.Delete(&models.Barang{}, id)

	c.JSON(200, gin.H{
		"message": "Berhasil menghapus barang",
	})
}
