package utils

import (
	"fmt"
	"net"
	"time"

	"github.com/google/uuid"
	"github.com/mojocn/base64Captcha"
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

// plugin generate uuid string
func GetUUID() string {
	return uuid.New().String()
}

// plugin generate safe string
func GenerateEncryptString(str string) string {
	b, _ := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)

	return string(b)
}

// plugin generate captcha
func GenerateCaptcha() (string, string, string) {
	driver := base64Captcha.NewDriverDigit(80, 240, 6, 1.0, 80)
	store := base64Captcha.DefaultMemStore
	captcha := base64Captcha.NewCaptcha(driver, store)

	id, image, answer, err := captcha.Generate()

	if err != nil {
		Logger.Error("generate captcha error")
	}

	return id, image, answer
}
