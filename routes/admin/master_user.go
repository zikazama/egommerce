package routes

import (
	"gitlab.com/zikazama/golang-final-project/config"
	"gitlab.com/zikazama/golang-final-project/helpers"
	"gitlab.com/zikazama/golang-final-project/models"

	"github.com/gin-gonic/gin"
)

func ReadUser(c *gin.Context) {
	id := c.Param("id")
	data := []models.User{}

	if id != "" {
		data := models.User{}
		if config.DB.Where("id = ?", id).Find(&data).RecordNotFound() {
			c.JSON(404, gin.H{
				"message": "User tidak ditemukan",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "User ditemukan",
			"data":    data,
		})
		return
	} else {
		config.DB.Find(&data)
		c.JSON(200, gin.H{
			"message": "User ditemukan",
			"data":    data,
		})
		return
	}
}

func StoreUser(c *gin.Context) {
	hashPassword, _ := helpers.HashPassword(c.PostForm("password"))
	data := models.User{
		Username: c.PostForm("username"),
		Email:    c.PostForm("email"),
		Password: hashPassword,
		Role:     "user",
		Nama:     c.PostForm("nama"),
		Alamat:   c.PostForm("alamat"),
	}

	config.DB.Create(&data)

	c.JSON(201, gin.H{
		"message": "Berhasil menambahkan pengirim",
		"data":    data,
	})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var data models.User
	if config.DB.Where("id = ?", id).Find(&data).RecordNotFound() {
		c.JSON(404, gin.H{
			"message": "User tidak ditemukan",
		})
		return
	}

	hashPassword, _ := helpers.HashPassword(c.PostForm("password"))
	dataUpdate := models.User{
		Username: c.PostForm("username"),
		Email:    c.PostForm("email"),
		Password: hashPassword,
		Role:     "user",
		Nama:     c.PostForm("nama"),
		Alamat:   c.PostForm("alamat"),
	}

	config.DB.Model(&data).Updates(dataUpdate)

	c.JSON(200, gin.H{
		"message": "Berhasil memperbarui user",
		"data":    data,
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var data models.User
	if config.DB.Where("id = ?", id).Find(&data).RecordNotFound() {
		c.JSON(404, gin.H{
			"message": "User tidak ditemukan",
		})
		return
	}

	config.DB.Delete(&models.User{}, id)

	c.JSON(200, gin.H{
		"message": "Berhasil menghapus user",
	})
}
