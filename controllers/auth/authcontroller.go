package auth

import (
	"errors"
	"net/http"
	"go-rest-with-gin/utils"
	user "go-rest-with-gin/models"
	"github.com/gin-gonic/gin"
	// "github.com/dgrijalva/jwt-go"
)


const UserCollection = "user"

var (

	errFindByUsername		=	errors.New("an occured during query")
	errInvalidUsername		=	errors.New("invalid username")
	errInvalidPassword		= 	errors.New("invalid password")
	errorInvalidBody		= 	errors.New("Invalid Request Body")
)

type LoginPayload struct {
	Username 	string `json:"username"`
	Password	string	`json:"password"`
}

func Login(c *gin.Context) {

	var loginpayload LoginPayload
	error := c.Bind(&loginpayload);
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "false", "message": errorInvalidBody.Error()})
		return
	}

	user, err := user.GetUserByUsername(loginpayload.Username, UserCollection)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "false", "message": errFindByUsername.Error()})
		return
	}

	if !utils.VerifyPassword(user.Password, []byte(loginpayload.Password)) {
		c.JSON(http.StatusBadRequest, gin.H{"status": "false", "message": errInvalidPassword.Error()})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "false", "message": errInvalidPassword.Error()})
		return
	}





}