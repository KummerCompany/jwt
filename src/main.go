package jwt

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/KummerCompany/jwt/src/config"
	"github.com/KummerCompany/jwt/src/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/valyala/fasthttp"
)

// CreateToken is function that return encrypt claims with a private key
func CreateToken(data interface{}) (string, error) {

	claims := models.Token{
		Data: data,
		Exp:  fmt.Sprint(time.Now().Add(time.Hour * 7 * 24).Unix()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(config.GetEnv("SECRET_TOKEN_AUTH")))

	if err != nil {
		return "", err
	}

	return t, nil
}

// ExtractToken for Authorization Bear
func extractToken(r *fasthttp.Request) string {
	bearToken := r.Header.Peek("Authorization")
	strArr := strings.Split(string(bearToken), " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

// MetadataToken extract bear information from header and transform en token struct
func MetadataToken(r *fasthttp.Request) (jwt.MapClaims, error) {
	tokenString := extractToken(r)

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.GetEnv("SECRET_TOKEN_AUTH")), nil
	})

	if err != nil || !token.Valid {
		return claims, err
	}

	timeToken, _ := strconv.ParseInt(claims["exp"].(string), 10, 64)

	if timeToken < time.Now().Unix() {
		return claims, fmt.Errorf("Token timed out")
	}

	return claims, err
}
