package ajwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
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

type TokenClaims struct {
	jwt.StandardClaims
	Login string
	Email string
}

func ParseToken(accessToken string, signingKey string) (TokenClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return TokenClaims{}, err
	}

	claims, ok := token.Claims.(TokenClaims)
	if !ok {
		return TokenClaims{}, errors.New("token claims are not of type *tokenClaims")
	}

	return claims, nil
}
