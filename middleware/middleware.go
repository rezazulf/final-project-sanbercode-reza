package middleware

import (
	"final-project/database"
	"final-project/models"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Auth(auth string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("Authorization")

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			c.Abort()
			return
		}

		// verify token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			c.Abort()
			return
		}

		// check token
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Check if the token has not expired
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "Token expired",
				})
				c.Abort()

				return
			}

			if claims["role"] != auth && auth != "isLogin" {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "Unauthorized",
				})
				c.Abort()

				return
			}

			db := database.DbConnection
			var user = models.Users{}

			//scan into db
			sql := "SELECT * FROM users WHERE id = $1"
			err := db.QueryRow(sql, claims["id"]).Scan(&user.ID, &user.Username, &user.Password, &user.Balance, &user.Role, &user.Created_at, &user.Updated_at)

			if err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
			}

			c.Set("user", user)

			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			c.Abort()
		}

	}
}
