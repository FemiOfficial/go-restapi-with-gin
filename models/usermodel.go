package models

import (
	"time"
	"go-rest-with-gin/conn"
	"gopkg.in/mgo.v2/bson"
)

// User structure
type User struct {
    ID        bson.ObjectId `bson:"_id"`
    Name      string        `bson:"name"`
    Username  string        `bson:"username"`
    Password  string        `bson:"password"`
    Address   string        `bson:"address"`
    Age       int           `bson:"age"`
    CreatedAt time.Time     `bson:"created_at"`
    UpdatedAt time.Time     `bson:"updated_at"`
}
// Users list
type Users []User
// GetUserById model function
func GetUserById(id bson.ObjectId, userCollection string) (User, error) {
    // Get DB from Mongo Config
    db := conn.GetMongoDB()
    user := User{}
    err := db.C(userCollection).Find(bson.M{"_id": &id}).One(&user)
    return user, err
}

// GetUserByUsername model function
func GetUserByUsername(username string, userCollection string) (User, error) {

    db := conn.GetMongoDB()
    user := User{}
    err := db.C(userCollection).Find(bson.M{"username": &username}).One(&user)
    return user, err

}