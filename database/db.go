package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ErrNoMatch = fmt.Errorf("no matching record")

var DB *gorm.DB

func Conn() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost, dbPort, dbUser, dbPassword, dbName := os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	database, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	// sqlDB, err := database.DB()
	// if err != nil {
	// 	// control error
	// 	panic("failed create pool")
	// }

	// sqlDB.SetMaxIdleConns(100)
	// // SetMaxOpenConns sets the maximum number of open connections to the database.
	// sqlDB.SetMaxOpenConns(1000)

	// // SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	// sqlDB.SetConnMaxLifetime(time.Hour)

	if err != nil {
		panic("Failed to connect to database!")
	} else {
		fmt.Println("Database connected!")
	}

	DB = database

}
