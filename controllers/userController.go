package controllers

import (
	"final-project/database"
	"final-project/models"
	"final-project/repository"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	db := database.DbConnection
	var body struct {
		Username string
		Password string
	}

	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Gagal read body JSON",
		})

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password tidak bisa di hash",
		})

		return
	}

	user := models.Users{Username: body.Username, Password: string(hash)}
	query := "INSERT INTO users (username, password, balance, role) VALUES ($1, $2, $3, $4)"
	err = db.QueryRow(query, user.Username, user.Password, "0", "Customer").Err()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Register Gagal!",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success to Create User",
	})

}

func Login(c *gin.Context) {
	db := database.DbConnection
	var body struct {
		Username string
		Password string
	}

	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Gagal read body JSON",
		})

		return
	}

	var users = models.Users{}

	//scan into db
	query := "SELECT id, username, password, role FROM users WHERE username = $1"
	err := db.QueryRow(query, body.Username).Scan(&users.ID, &users.Username, &users.Password, &users.Role)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Username atau Password salah",
		})
		panic(err)
	}

	//compare password with db password
	err = bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Username atau Password salah",
		})
		panic(err)
	}

	//make new jwt claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   users.ID,
		"role": users.Role,
		"exp":  time.Now().Add(time.Hour * 12).Unix(),
	})

	//make new token
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Gagal Create Token!",
		})
		panic(err)
	}

	//make cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*12, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})

}

func Logout(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
		return
	}

	c.SetCookie("Authorization", tokenString, -1, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil logout",
	})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"message": "Masuk sebagai",
		"user":    user,
	})
}

func NewPassword(c *gin.Context) {
	db := database.DbConnection
	var body struct {
		Password string
	}

	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Gagal read body JSOn",
		})

		return
	}

	user, _ := c.Get("user")
	ok := bcrypt.CompareHashAndPassword([]byte(user.(models.Users).Password), []byte(body.Password))
	if ok == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password sama dengan password lama",
		})
		c.Abort()
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password tidak bisa di hash",
		})

		return
	}

	query := "UPDATE users SET password = $1, updated_at = $2 WHERE id = $3"

	err = db.QueryRow(query, hash, time.Now().Format(time.RFC3339), user.(models.Users).ID).Err()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Gagal mengganti password",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Password telah diganti",
	})
}

func GetAdmins(c *gin.Context) {
	var (
		result gin.H
	)

	admin, err := repository.GetAdmins(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": admin,
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetUsers(c *gin.Context) {
	var (
		result gin.H
	)

	user, err := repository.GetUser(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": user,
		}
	}

	c.JSON(http.StatusOK, result)
}

func EditBalance(c *gin.Context) {
	var users models.Users

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&users)
	if err != nil {
		panic(err)
	}

	users.ID = int64(id)

	err = repository.EditBalance(database.DbConnection, &users)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Balance Updated",
	})
}
