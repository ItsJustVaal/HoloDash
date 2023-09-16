package database

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDatase() *sql.DB {
	var envs map[string]string
	envs, dberr := godotenv.Read(".env")

	if dberr != nil {
		log.Fatal("Error loading .env file")
	}

	// converting port to int
	port, _ := strconv.Atoi(envs["DBPORT"])

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", envs["DBHOST"], port, envs["DBUSER"], envs["DBPASSWORD"], envs["DBNAME"])

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal(err)
	}

	// check db
	err = db.Ping()
	if dberr != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected!")
	return db
}
