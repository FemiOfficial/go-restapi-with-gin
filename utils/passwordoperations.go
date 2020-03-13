package utils

import (
	"log"
  "golang.org/x/crypto/bcrypt"
)

func ConvertStrToByte(str string) []byte {
	return []byte(str)
}

func HashPassword(pwd []byte) string {
  hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost);
  if err != nil {
    log.Fatal(err)
  }
  return string(hash)
}

func VerifyPassword(pwdhash string, plainpassword []byte) bool {
  bytehash := []byte(pwdhash);
  err := bcrypt.CompareHashAndPassword(bytehash, plainpassword)
  if err != nil {
    // log.Fatal(err)
    return false
  }

  return true
}
