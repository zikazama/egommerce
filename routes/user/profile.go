package routes

import (
	"gitlab.com/zikazama/golang-final-project/config"
	"gitlab.com/zikazama/golang-final-project/helpers"
	"gitlab.com/zikazama/golang-final-project/models"

	"github.com/gin-gonic/gin"
)

func ReadProfile(c *gin.Context) {
	user_id, _ := c.Get("user_id")
	parse_user_id := user_id.(float64)
	parse_user_id_int := int(parse_user_id)

	data := models.User{}
	if config.DB.Where("id = ?", parse_user_id_int).Find(&data).RecordNotFound() {
		c.JSON(404, gin.H{
			"message": "Profile tidak ditemukan",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Profile ditemukan",
		"data":    data,
	})
}

func UpdateProfile(c *gin.Context) {
	user_id, _ := c.Get("user_id")
	parse_user_id := user_id.(float64)
	parse_user_id_int := int(parse_user_id)
	var data models.User
	if config.DB.Where("id = ?", parse_user_id_int).Find(&data).RecordNotFound() {
		c.JSON(404, gin.H{
			"message": "Profile tidak ditemukan",
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
		"message": "Berhasil memperbarui profile",
		"data":    data,
	})
}
