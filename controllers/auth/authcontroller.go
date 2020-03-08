package auth

import (
	"errors"
	"net/http"
	"time"
	"go-rest-with-gin/conn"
	user "go-rest-with-gin/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"github.com/dgrijalva/jwt-go"
)


const UserCollection = "user"

type struct LoginPayload {
	Username 	string `json:"username"`
	Password	string	`json:password`
}

var (

	errFindByUsername		=	errors.New("an occured during query")
	errInvalidUsername		=	errors.New("invalid username")
	errInvalidPassword		= 	errors.New("invalid password")
	errorInvalidBody		= 	errors.New("Invalid Request Body")
)

func Login(c *gin.Context) {

	db := conn.GetMongoDB()

	var loginpayload LoginPayload
	error := c.Bind(&loginpayload);
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "false", "message": errorInvalidBody.Error()})
		return
	}

	user, err = user.GetUserByUsername(loginpayload.username, )

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "false", "message": errFindByUsername.Error()})
		return
	}

	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "false", "message": errInvalidUsername.Error()})
		return
	}





}