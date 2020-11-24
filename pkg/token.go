package pkg

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pass []byte) string {
	hashed, err := bcrypt.Cost(pass)
	if err != nil {
		log.Print("Bcrypt pass error: ", err)
		return ""
	}
	return string(hashed)
}
