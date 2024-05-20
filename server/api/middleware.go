package api

import (
	"net/http"
	"tournament_api/server/utils"
)

const API_KEY = "X-Api-Key"
const EXPECTED_API_KEY = "tournament"

func (s *Server) HeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		w.Header().Set("Content-Security-Policy", "default-src 'self';base-uri 'self';font-src 'self' https: data:;form-action 'self';frame-ancestors 'self';img-src 'self' data:;object-src 'none';script-src 'self';script-src-attr 'none';style-src 'self' https: 'unsafe-inline';upgrade-insecure-requests")
		w.Header().Set("Cross-Origin-Opener-Policy", "same-origin")
		w.Header().Set("Cross-Origin-Resource-Policy", "same-origin")
		w.Header().Set("Origin-Agent-Cluster", "?1")
		w.Header().Set("Referrer-Policy", "no-referrer")
		w.Header().Set("Strict-Transport-Security", "max-age=15552000; includeSubDomains")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-DNS-Prefetch-Control", "off")
		w.Header().Set("X-Download-Options", "noopen")
		w.Header().Set("X-Frame-Options", "SAMEORIGIN")
		w.Header().Set("X-Permitted-Cross-Domain-Policies", "none")
		w.Header().Set("X-XSS-Protection", "0")

		next.ServeHTTP(w, r)
	})
}

func (s *Server) AuthenticateMiddleware(next http.Handler) http.Handler {
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
func (s *Server) CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		next.ServeHTTP(w, r)
	})
}
