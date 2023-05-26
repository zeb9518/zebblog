package utils

import (
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Issuer string
	jwt.MapClaims
}

/**
 * @Description: 生成token
 * @param {string} issuer 生成者
 * @param {int64} expiresAt 过期时间
 */
func GenerateToke(secretKey string, issuer string, expiresAt int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["Issuer"] = issuer
	claims["ExpiresAt"] = expiresAt
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

/**
 * @Description: 验证Token
 * @param {string} token	待验证的Token
 */
func VerifyToken(secretKey string, token string) error {
	_, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	return err
}
