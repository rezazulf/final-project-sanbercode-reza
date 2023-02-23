package main

import (
	"database/sql"
	"final-project/database"
	"final-project/routers"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load file env")
	} else {
		fmt.Println("successfully read file env")
	}
	// DB Test
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	DB, err = sql.Open("postgres", psqlInfo)
	err := DB.Ping()
	if err != nil {
		fmt.Println("DB Conn Failed")
	} else {
		fmt.Println("DB Conn Succeeded")
	}

	database.DbMigrate(DB)

	defer DB.Close()

	routers.StartServer().Run("localhost:8080")
}