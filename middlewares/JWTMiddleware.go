package middlewares

import (
	"fmt"
	"strings"

	"gitlab.com/zikazama/golang-final-project/config"
	"gitlab.com/zikazama/golang-final-project/helpers"
	"gitlab.com/zikazama/golang-final-project/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func IsAuth(role string) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.Request.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) < 2 {
			c.JSON(200, gin.H{
				"message": "Authorization tidak ditemukan",
			})
			c.Abort()
			return
		}

		token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(helpers.SecretJWT), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims["user_id"], claims["name"], claims["role"])
			c.Set("user_id", claims["user_id"])
			c.Set("name", claims["name"])
			c.Set("role", claims["role"])
			var data models.User
			if config.DB.Where("id = ? AND role = ?", claims["user_id"], claims["role"]).Find(&data).RecordNotFound() {
				c.JSON(404, gin.H{
					"message": "User dari token tidak ditemukan",
				})
				c.Abort()
				return
			}
			if claims["role"] != role {
				c.JSON(401, gin.H{
					"message": "Akses tidak diizinkan",
				})
				c.Abort()
				return
			}
		} else {
			fmt.Println(err)
			c.JSON(402, gin.H{
				"message": "Invalid token",
			})
			c.Abort()
			return
		}
	}
}
