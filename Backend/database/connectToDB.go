package database

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDatase() {
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
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
