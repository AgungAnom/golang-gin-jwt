package helpers

import "github.com/dgrijalva/jwt-go"

var secretKey = "secret"

func GenerateToken(id uint, email string) (token string, err error) {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = parseToken.SignedString([]byte(secretKey))

	return
}