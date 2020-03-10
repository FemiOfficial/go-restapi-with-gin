package utils

import (
	"log"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func convertStrToByte(str string) []byte {
	return []byte(str)
}

func hashPassword(pwd []byte) string {
  hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost);
  if err != nil {
    log.Fatal(err)
  }
  return string(hash)
}

func verifyPassword(pwdhash string, plainpassword []byte) bool {
  bytehash := []byte(pwdhash);
  err := bcrypt.CompareHashAndPassword(bytehash, plainpassword)
  if err != nil {
    // log.Fatal(err)
    return false
  }

  return true
}
