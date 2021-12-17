package ajwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(d time.Duration, signingKey []byte, cl map[string]interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	for key, val := range cl {
		claims[key] = val
	}

	claims["authorized"] = true
	claims["exp"] = time.Now().Add(d).Unix()

	tokenString, err := token.SignedString(signingKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
