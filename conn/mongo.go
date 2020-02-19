package conn

import (
	"fmt"
	"os"
	"gopkg.in/mgo.v2"
)

var db *mgo.Database
func init() {
	host := os.Getenv("MONGO_HOST")
	dbName := os.Getenv("MONGO_DB_NAME")
	session, err := mgo.Dial(host)
	if err != nil {
		fmt.Println("DB session err: ", err)
		os.Exit(2)
	}
	db = session.DB(dbName)
}

func GetMongoDB() *mgo.Database {
	return db
}