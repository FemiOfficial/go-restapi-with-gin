package users

import (
	"errors"
	"net/http"
	"time"
	// "log"
	"go-rest-with-gin/config"
	"go-rest-with-gin/utils"
	user "go-rest-with-gin/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	// "gopkg.in/validator.v2"
)

const UserCollection = "user"

var (
	errNotExist			= errors.New("User does not exist")
	errorInvalidId		= errors.New("User with Id Does not Exist")
	errorInvalidBody	= errors.New("Invalid Request Body")
	errInsertionFailed	= errors.New("Error in DB Insertion for New User")
	errUpdateFailed		= errors.New("Error in user update")
	errDeleteFailed 	= errors.New("Error while deleteng User")
)


func GetAllUsers(c *gin.Context) {
	db := config.GetMongoDB()
	users := user.Users{}
	err := db.C(UserCollection).Find(bson.M{}).All(&users)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errNotExist.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "users": &users})
}

func GetUserById(c *gin.Context) {
	var id bson.ObjectId = bson.ObjectIdHex(c.Param("id"))
	user, err := user.GetUserById(id, UserCollection)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errorInvalidId.Error()})
		return 
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "user": &user})
}

func RegisterUser(c *gin.Context) {
	// Get DB from Mongo Config
	db := config.GetMongoDB()
	user := user.User{}

	if error := c.ShouldBind(&user); error != nil {
		// log.Fatal(error)
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed",
		"message": errorInvalidBody.Error()})
		return
	}

	user.ID = bson.NewObjectId()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	bytepassword := utils.ConvertStrToByte(user.Password)
	user.Password = utils.HashPassword(bytepassword)


	err := db.C(UserCollection).Insert(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInsertionFailed.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "user": &user})
	return
}

func UpdateUser(c *gin.Context) {
	db := config.GetMongoDB()
	var id bson.ObjectId = bson.ObjectIdHex(c.Param("id"))
	existingUser, err := user.GetUserById(id, UserCollection)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errorInvalidId.Error()})
		return
	}

	err = c.Bind(&existingUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errorInvalidBody.Error()})
		return
	}

	existingUser.ID = id
	existingUser.UpdatedAt = time.Now()
	err = db.C(UserCollection).Update(bson.M{"_id": id}, existingUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "success", "message": errUpdateFailed.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": &existingUser})
}

func DeleteUser(c *gin.Context) {
	db := config.GetMongoDB()
	var id bson.ObjectId = bson.ObjectIdHex(c.Param("id"))
	err := db.C(UserCollection).Remove(bson.M{"_id": &id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errDeleteFailed.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "user deleted successfully"})

}
