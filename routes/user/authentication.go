package routes

import (
	"net/http"
	"time"

	"gitlab.com/zikazama/golang-final-project/config"
	"gitlab.com/zikazama/golang-final-project/helpers"
	"gitlab.com/zikazama/golang-final-project/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	hashPassword, _ := helpers.HashPassword(c.PostForm("password"))
	data := models.User{
		Email:    c.PostForm("email"),
		Password: hashPassword,
		Username: c.PostForm("username"),
		Role:     "user",
		Nama:     c.PostForm("nama"),
		Alamat:   c.PostForm("alamat"),
	}

	config.DB.Create(&data)

	c.JSON(200, gin.H{
		"message": "Registrasi Berhasil",
		"data":    data,
	})
}

func Login(c *gin.Context) {
	var data models.User

	email := c.PostForm("email")
	password := c.PostForm("password")

	if config.DB.Where("email = ? and role = ?", email, "user").Find(&data).RecordNotFound() {
		c.JSON(200, gin.H{
			"message": "Email salah",
		})
		return
	}

	if !helpers.CheckPasswordHash(password, data.Password) {
		c.JSON(200, gin.H{
			"message": "Kata sandi salah",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": data.ID,
		"name":    data.Nama,
		"role":    data.Role,
		"exp":     time.Now().AddDate(0, 0, 7).Unix(),
		"iat":     time.Now().Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(helpers.SecretJWT))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"message": "User ditemukan",
		"token":   tokenString,
	})
}
