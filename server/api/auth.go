package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	FirstName string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	jwt.RegisteredClaims
}

const ACCESS_TOKEN = "access_token"
const REFRESH_TOKEN = "refresh_token"

type AuthToken struct {
	name     string
	key      []byte
	duration time.Time
}

func (server *Server) newAccessToken() (t *AuthToken) {
	accessToken := &AuthToken{
		name:     ACCESS_TOKEN,
		key:      []byte(server.config.JWT_ACCESS_SECRET),
		duration: time.Now().Add(time.Second * time.Duration(server.config.JWT_ACCESS_EXPIRATION_IN_SECONDS)),
	}
	return accessToken
}

func (server *Server) newRefreshToken() (t *AuthToken) {
	refreshToken := &AuthToken{
		name:     REFRESH_TOKEN,
		key:      []byte(server.config.JWT_REFRESH_SECRET),
		duration: time.Now().Add(time.Second * time.Duration(server.config.JWT_REFRESH_EXPIRATION_IN_SECONDS)),
	}
	return refreshToken
}

func (t *AuthToken) generateTokenString(Email *string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		Email: *Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(t.duration),
		},
	})
	return token.SignedString(t.key)
}

func (t *AuthToken) verifyToken(tokenStr string) (*Claims, error) {
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

func (t *AuthToken) saveTokenAsCookie(w http.ResponseWriter, value string) {
	http.SetCookie(w, &http.Cookie{
		Name:     t.name,
		Value:    value,
		Expires:  t.duration,
		HttpOnly: true,
		Path:     "/",
	})
}
