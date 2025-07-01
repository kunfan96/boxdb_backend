package utils

import (
	"crypto/rand"
	"encoding/hex"
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
func GeUUID() string {
	return uuid.New().String()
}

func GenerateToken(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// plugin generate safe string
func GenerateEncryptString(str string) string {
	b, _ := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)

	return string(b)
}

// 校验密码
func CheckPassword(passwordInSql, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(passwordInSql), []byte(password)) == nil
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
