package main

import (
	"log"
	"os"

	"github.com/go-pg/pg/v9"
	"github.com/joho/godotenv"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
)

const directory = "cmd/migrations"

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Could not load .env file: %s", err.Error())
	}

	db := pg.Connect(&pg.Options{
		Addr:     os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Database: os.Getenv("DB_DATABASE"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	err = migrations.Run(db, directory, os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
