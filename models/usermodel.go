package usermodel

import (
    "fmt"
	"time"
	"go-rest-with-gin/config"
    "gopkg.in/mgo.v2/bson"
)

// User structure
type User struct {
    ID        bson.ObjectId `bson:"_id"`
    Name      string        `bson:"name" form:"name" json:"name" binding:"required"`
    Username  string        `bson:"username" form:"username" json:"username" binding:"required"`
    Password  string        `bson:"password" form:"password" json:"password" binding:"required"`
    Address   string        `bson:"address" form:"address" json:"address" binding:"required"`
    Age       int           `bson:"age" form:"age" json:"age" binding:"required"`
    CreatedAt time.Time     `bson:"created_at"`
    UpdatedAt time.Time     `bson:"updated_at"`
}
// Users list
type Users []User
// GetUserById model function
func GetUserById(id bson.ObjectId, userCollection string) (User, error) {
    // Get DB from Mongo Config
    db := config.GetMongoDB()
    user := User{}
    err := db.C(userCollection).Find(bson.M{"_id": &id}).One(&user)
    return user, err
}

// GetUserByUsername model function
func GetUserByUsername(username string, userCollection string) (User, error) {
    fmt.Println(username)
    db := config.GetMongoDB()
    user := User{}
    err := db.C(userCollection).Find(bson.M{"username": &username}).One(&user)
    return user, err

}