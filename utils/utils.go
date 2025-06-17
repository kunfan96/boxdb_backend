package utils

import (
	"fmt"
	"net"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// check port is used
func IsPortUsed(port string) bool {
	addr := fmt.Sprintf(":%s", port)
	conn, err := net.DialTimeout("tcp", addr, time.Second*3)

	// connect success
	if err == nil {
		// close connect
		conn.Close()

		return true
	}

	return false
}

func EncryptString(str string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)

	return string(b), err
}

func GetUUID() string {
	return uuid.New().String()
}
