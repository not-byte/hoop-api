package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Username string `json:"username"`
	Lastname string `json:"lastname"`
	jwt.RegisteredClaims
}
type TOKEN_TYPE string

const ACCESS_TOKEN = "access_token"
const REFRESH_TOKEN = "refresh_token"

type Token struct {
	name     string
	key      []byte
	duration time.Time
}

func newAccessToken() (t *Token) {
	accessToken := &Token{
		name:     ACCESS_TOKEN,
		key:      []byte("secret"),
		duration: time.Now().Add(time.Minute * 10),
	}
	return accessToken
}

func newRefreshToken() (t *Token) {
	refreshToken := &Token{
		name:     REFRESH_TOKEN,
		key:      []byte("secret"),
		duration: time.Now().Add(time.Hour * 24 * 15),
	}
	return refreshToken
}

func (t *Token) generateTokenString(FirstName *string, LastName *string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		Username: *FirstName,
		Lastname: *LastName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(t.duration),
		},
	})
	return token.SignedString(t.key)
}

func (t *Token) verifyToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return t.key, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}

func (t *Token) saveTokenAsCookie(w http.ResponseWriter, value string) {
	http.SetCookie(w, &http.Cookie{
		Name:     t.name,
		Value:    value,
		Expires:  t.duration,
		HttpOnly: true,
		Path:     "/",
	})
}
