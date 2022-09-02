package auth

import (
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"strings"
	"time"
)

var signedKey = []byte("golang_signed_key")
var tokenHead = "Bearer "

// 密码验证
func CheckPassword(raw, pwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(raw), []byte(pwd))
}

// 生成token
func GenerateToken(id int64, name string) map[string]string {
	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   strconv.Itoa(int(id)),
		"name": name,
		"exp":  time.Now().Unix() + 3600,
	}).SignedString(signedKey)
	if err != nil {
		return nil
	}
	return map[string]string{
		"tokenHead": tokenHead,
		"token":     tokenString,
	}
}

func GetToken(tokenString string) string {
	if !strings.Contains(tokenString, tokenHead) {
		return ""
	}
	tokenString = strings.Replace(tokenString, tokenHead, "", 1)
	return tokenString
}

// 解析token
func ParseToken(tokenString string) (string, string) {
	if tokenString = GetToken(tokenString); tokenString == "" {
		return "", ""
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signedKey, nil
	})
	if err != nil {
		return "", ""
	}
	claims := token.Claims.(jwt.MapClaims)
	return claims["id"].(string), claims["name"].(string)
}

// 刷新token
func RefreshToken(tokenString string) map[string]string {
	if tokenString = GetToken(tokenString); tokenString == "" {
		return nil
	}
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signedKey, nil
	})
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Unix() + 3600
	refreshToken, err := token.SignedString(signedKey)
	if err != nil {
		return nil
	}
	return map[string]string{
		"tokenHead": tokenHead,
		"token":     refreshToken,
	}
}
