package utils

import (
	// "log"
	"os"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"go-rest-with-gin/config"
)

type AuthDetails struct {
  Name		string			`json:"name"`
  Username	string			`json:"username"`
  Address	string			`json:"address"`
  Age		int				`json:"age"`
}

func CreateToken(userDetails AuthDetails) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorirzed"] = true
	claims["name"] = userDetails.Name
	claims["username"] = userDetails.Username
	claims["address"] = userDetails.Address
	claims["age"] = userDetails.Age
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Getenv("API_SECRET")))
}

func IsTokenValid(authorization string) error {
	token, err := verifyToken(authorization)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//does this token conform to "SigningMethodHMAC" ?
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func GetTokenData(tokenString string) (*AuthDetails, error) {
	token, err := verifyToken(tokenString)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims

	if ok && token.Valid {

		name, ok := claims["name"].(string) //convert the interface to string
		if !ok {
			return nil, err
		}
		username, ok := claims["username"].(string) //convert the interface to string
		if !ok {
			return nil, err
		}

		address, ok := claims["address"].(string) //convert the interface to string
		if !ok {
			return nil, err
		}

		age, ok := claims["age"].(int) //convert the interface to string
		if !ok {
			return nil, err
		}

		return &AuthDetails{
			Name:   	name,
			Username:   username,
			Address:   	address,
			Age:   		age,
		}, nil
	}
	return nil, err
}