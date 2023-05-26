package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// 生成随机的 salt 字符串
func generateSalt(length int) string {
	salt := make([]rune, length)
	for i := range salt {
		salt[i] = letters[rand.Intn(len(letters))]
	}
	return string(salt)
}

// 对密码进行加盐和 MD5 加密
func EncryptPassword(password string) (string, string) {
	salt := generateSalt(10)
	saltedPassword := password + salt
	hasher := md5.New()
	hasher.Write([]byte(saltedPassword))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))
	return hashedPassword, salt
}

// 验证密码是否正确
func ValidatePassword(password, hashedPassword, salt string) bool {
	saltedPassword := password + salt
	hasher := md5.New()
	hasher.Write([]byte(saltedPassword))
	return hashedPassword == hex.EncodeToString(hasher.Sum(nil))
}
