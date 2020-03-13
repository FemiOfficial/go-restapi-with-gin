package config

import (
	"fmt"
	"os"
	"log"
	"gopkg.in/mgo.v2"
	"github.com/joho/godotenv"
)

var db *mgo.Database

func Getenv(key string) string{
  err := godotenv.Load(".env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }

  return os.Getenv(key)
}

func GetMongoDB() *mgo.Database {
	host := Getenv("MONGO_HOST")
	dbName := Getenv("MONGO_DB_NAME")

	session, err := mgo.Dial(host)
	if err != nil {
		fmt.Println("DB session err: ", err)
		os.Exit(2)
	}
	db = session.DB(dbName)
	return db
}