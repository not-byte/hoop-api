package api

import (
	"net/http"
	"tournament_api/server/utils"
)

const API_KEY = "API KEY"
const EXPECTED_API_KEY = "MY API KEY"

func (s *Server) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessTokenString, accessTokenErr := r.Cookie(ACCESS_TOKEN)
		refreshTokenString, refreshTokenErr := r.Cookie(REFRESH_TOKEN)

		if accessTokenErr != nil {
			utils.HandleCookieError(w, accessTokenErr, ACCESS_TOKEN)
			return
		}

		if refreshTokenErr != nil {
			utils.HandleCookieError(w, refreshTokenErr, REFRESH_TOKEN)
			return
		}

		accessToken := s.newAccessToken()
		refreshToken := s.newRefreshToken()

		if _, err := accessToken.verifyToken(accessTokenString.Value); err != nil {
			http.Error(w, "Invalid access token", http.StatusUnauthorized)
			return
		}

		if _, err := refreshToken.verifyToken(refreshTokenString.Value); err != nil {
			http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (s *Server) TokenRefreshMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessTokenString, accessTokenErr := r.Cookie(ACCESS_TOKEN)
		refreshTokenString, refreshTokenErr := r.Cookie(REFRESH_TOKEN)

		if refreshTokenErr != nil || refreshTokenString == nil || refreshTokenString.Value == "" {
			http.Error(w, "Unauthorized - Refresh Token required", http.StatusUnauthorized)
			return
		}

		refreshToken := s.newRefreshToken()

		refreshTokenClaims, refErr := refreshToken.verifyToken(refreshTokenString.Value)
		if refErr != nil {
			http.Error(w, "Unauthorized - Invalid Refresh Token", http.StatusUnauthorized)
			return
		}

		if accessTokenString == nil || accessTokenErr == http.ErrNoCookie {
			newAccessToken := s.newAccessToken()

			email := refreshTokenClaims.Email

			newAccessTokenString, err := newAccessToken.generateTokenString(&email)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			newAccessToken.saveTokenAsCookie(w, newAccessTokenString)
		}
		next.ServeHTTP(w, r)
	})
}

func (s *Server) APIKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get(API_KEY)

		if apiKey != EXPECTED_API_KEY {
			http.Error(w, "Unauthorized: Missing or incorrect API key", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// this will only run on development
func (s *Server) CORSmiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		next.ServeHTTP(w, r)
	})
}
