package auth

import (
	"errors"
	"net/http"
	"log"
	"go-rest-with-gin/utils"
	usermodel "go-rest-with-gin/models"
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
	Username 	string `form:"username" json:"username" binding:"required"`
	Password	string	`form:"password" json:"password" binding:"required"`
}

// type AuthDetails struct {
//   Name		string			`json:"name"`
//   Username	string			`json:"username"`
//   Address	string			`json:"address"`
//   Age		int				`json:"age"`
// }


func Login(c *gin.Context) {

	var loginpayload LoginPayload

	error := c.ShouldBind(&loginpayload);

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "false", "message": errorInvalidBody.Error()})
		return
	}

	user, err := usermodel.GetUserByUsername(loginpayload.Username, UserCollection)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "false", "message": errInvalidUsername.Error()})
		return
	}

	if !utils.VerifyPassword(user.Password, []byte(loginpayload.Password)) {
		c.JSON(http.StatusBadRequest, gin.H{"status": "false", "message": errInvalidPassword.Error()})
		return
	} else {

		tokenPayload := utils.AuthDetails{ user.Name, user.Username, user.Address, user.Age }
		token, err := utils.CreateToken(tokenPayload)

		if(err != nil) {
			log.Fatal(err)
		}

		// log.Fatal(token)

		c.JSON(http.StatusOK, gin.H{"status": "false",
		"message": "success login","token": token, "data": tokenPayload})
		return
	}


}