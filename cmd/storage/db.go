package storage

import (
	"database/sql"
	"log"
	"os"
	"fmt"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

)

var db *sql.DB

func InitDB(){
	err := godotenv.Load()
	if err !=nil {
		log.Fatal("Error loading env variables")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	db, err = sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort))
    if err != nil {
		panic(err.Error())
	} 

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	} 

	fmt.Println("DB Connection was successful")
}


func GetDB() * sql.DB{
	return db
}