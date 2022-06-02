package routes

import (
	"gitlab.com/zikazama/golang-final-project/config"
	"gitlab.com/zikazama/golang-final-project/models"

	"github.com/gin-gonic/gin"
)

func ReadKategori(c *gin.Context) {
	id := c.Param("id")
	data := []models.Kategori{}

	if id != "" {
		data := models.Kategori{}
		if config.DB.Where("id = ?", id).Find(&data).RecordNotFound() {
			c.JSON(404, gin.H{
				"message": "Kategori tidak ditemukan",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Kategori ditemukan",
			"data":    data,
		})
		return
	} else {
		config.DB.Find(&data)
		c.JSON(200, gin.H{
			"message": "Kategori ditemukan",
			"data":    data,
		})
		return
	}
}

func StoreKategori(c *gin.Context) {
	data := models.Kategori{
		Nama_kategori: c.PostForm("nama_kategori"),
	}

	config.DB.Create(&data)

	c.JSON(201, gin.H{
		"message": "Berhasil menambahkan kategori",
		"data":    data,
	})
}

func UpdateKategori(c *gin.Context) {
	id := c.Param("id")
	var data models.Kategori
	if config.DB.Where("id = ?", id).Find(&data).RecordNotFound() {
		c.JSON(404, gin.H{
			"message": "Kategori tidak ditemukan",
		})
		return
	}

	dataUpdate := models.Kategori{
		Nama_kategori: c.PostForm("nama_kategori"),
	}

	config.DB.Model(&data).Updates(dataUpdate)

	c.JSON(200, gin.H{
		"message": "Berhasil memperbarui kategori",
		"data":    data,
	})
}

func DeleteKategori(c *gin.Context) {
	id := c.Param("id")
	var data models.Kategori
	if config.DB.Where("id = ?", id).Find(&data).RecordNotFound() {
		c.JSON(404, gin.H{
			"message": "Kategori tidak ditemukan",
		})
		return
	}

	config.DB.Delete(&models.Kategori{}, id)

	c.JSON(200, gin.H{
		"message": "Berhasil menghapus kategori",
	})
}
