package routes

import (
	"gitlab.com/zikazama/golang-final-project/config"
	"gitlab.com/zikazama/golang-final-project/models"

	"github.com/gin-gonic/gin"
)

func ReadPengirim(c *gin.Context) {
	id := c.Param("id")
	data := []models.Pengirim{}

	if id != "" {
		data := models.Pengirim{}
		if config.DB.Where("id = ?", id).Find(&data).RecordNotFound() {
			c.JSON(404, gin.H{
				"message": "Pengirim tidak ditemukan",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Pengirim ditemukan",
			"data":    data,
		})
		return
	} else {
		config.DB.Find(&data)
		c.JSON(200, gin.H{
			"message": "Pengirim ditemukan",
			"data":    data,
		})
		return
	}
}

func StorePengirim(c *gin.Context) {
	data := models.Pengirim{
		Nama_pengirim: c.PostForm("nama_pengirim"),
	}

	config.DB.Create(&data)

	c.JSON(201, gin.H{
		"message": "Berhasil menambahkan pengirim",
		"data":    data,
	})
}

func UpdatePengirim(c *gin.Context) {
	id := c.Param("id")
	var data models.Pengirim
	if config.DB.Where("id = ?", id).Find(&data).RecordNotFound() {
		c.JSON(404, gin.H{
			"message": "Pengirim tidak ditemukan",
		})
		return
	}

	dataUpdate := models.Pengirim{
		Nama_pengirim: c.PostForm("nama_pengirim"),
	}

	config.DB.Model(&data).Updates(dataUpdate)

	c.JSON(200, gin.H{
		"message": "Berhasil memperbarui pengirim",
		"data":    data,
	})
}

func DeletePengirim(c *gin.Context) {
	id := c.Param("id")
	var data models.Pengirim
	if config.DB.Where("id = ?", id).Find(&data).RecordNotFound() {
		c.JSON(404, gin.H{
			"message": "Pengirim tidak ditemukan",
		})
		return
	}

	config.DB.Delete(&models.Pengirim{}, id)

	c.JSON(200, gin.H{
		"message": "Berhasil menghapus pengirim",
	})
}
