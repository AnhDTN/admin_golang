package pkg

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pass []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	if err != nil {
		log.Print("Bcrypt pass error_custom: ", err)
		return ""
	}
	return string(hashed)
}
