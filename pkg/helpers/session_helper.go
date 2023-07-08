package helpers

import (
	"os"
	"time"

	"github.com/devhijazi/go-users-api/pkg/errors"
	"github.com/dgrijalva/jwt-go"
)

type JWTClaims struct {
	ID string `json:"id"`

	jwt.StandardClaims
}

var jwtToken = []byte(os.Getenv("JWT_TOKEN"))

func GenerateSessionToken(id string) string {

	claims := &JWTClaims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	str, _ := token.SignedString(jwtToken)

	return str
}

func ValidateAndDecodeSessionToken(signedToken string) (*JWTClaims, *errors.Error) {

	token, err := jwt.ParseWithClaims(signedToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) { return []byte(jwtToken), nil })

	if err != nil {
		return nil, errors.APIError()
	}

	claims, ok := token.Claims.(*JWTClaims)

	if !ok {
		return nil, errors.TokenError()
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return nil, errors.TokenError()
	}

	return claims, nil

}
