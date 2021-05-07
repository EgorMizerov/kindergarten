package hash

import (
	"crypto/sha256"
	"fmt"
)

func PasswordHash(password string) string {
	h := sha256.New()
	h.Write([]byte("password"))

	return fmt.Sprintf("%x", h.Sum(nil))
}
