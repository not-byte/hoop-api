package api

import (
	"encoding/json"
	"net/http"
	"tournament_api/server/model"
	"tournament_api/server/types"
	"tournament_api/server/utils"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

/*
think about it
type AuthHandlers struct {
	store interface{}
}

func (s *Server) authHandlers() *AuthHandlers {
	return &AuthHandlers{
		store: s.store,
	}
}
*/

func (s *Server) handleGetAll(w http.ResponseWriter, r *http.Request) {
	account, err := s.store.GetAccountByEmail("pawellinek2@gmail.com")
	if err != nil {
		http.Error(w, "Invalid login credentials "+err.Error(), http.StatusUnauthorized)
		return
	}
	if account == nil {
		http.Error(w, "Account does not exist", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(*account)
}

func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) {

	var user model.Account
	errDecode := json.NewDecoder(r.Body).Decode(&user)
	if errDecode != nil {
		http.Error(w, "Error while decoding: ", http.StatusInternalServerError)
		return
	}

	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		http.Error(w, "Validation failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	account, err := s.store.GetAccountByEmail(user.Email)
	if err != nil {
		http.Error(w, "Error searching for account", http.StatusInternalServerError)
		return
	}
	if account == nil {
		http.Error(w, "Account does not exist", http.StatusBadRequest)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(user.Password))
	if err != nil {
		http.Error(w, "Invalid login credentials", http.StatusUnauthorized)
		return
	}

	accessToken := s.newAccessToken()
	refreshToken := s.newRefreshToken()

	accessTokenString, errAccess := accessToken.generateTokenString(&user.Email)
	refreshTokenString, errRefresh := refreshToken.generateTokenString(&user.Email)

	if errAccess != nil || errRefresh != nil {
		http.Error(w, "Error generating token string", http.StatusInternalServerError)
		return
	}

	accessToken.saveTokenAsCookie(w, accessTokenString)
	refreshToken.saveTokenAsCookie(w, refreshTokenString)

	if err := json.NewEncoder(w).Encode(map[string]string{"message": "login successful"}); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func (s *Server) handleRegister(w http.ResponseWriter, r *http.Request) {

	var user types.User
	errDecode := json.NewDecoder(r.Body).Decode(&user)
	if errDecode != nil {
		http.Error(w, "Error while decoding: ", http.StatusInternalServerError)
		return
	}

	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		http.Error(w, "Validation failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	account, err := s.store.GetAccountByEmail(*user.Email)
	if account != nil {
		http.Error(w, "Account alredy exists "+err.Error(), http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, "Error "+err.Error(), http.StatusInternalServerError)
		return
	}

	hashedPassword, err := utils.HashPassword(*user.Password)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = s.store.CreateAccount(r.Context(), *user.Email, hashedPassword, 9)
	if err != nil {
		http.Error(w, "Error while creating account", http.StatusInternalServerError)
		return
	}

	accessToken := s.newAccessToken()
	refreshToken := s.newRefreshToken()

	accessTokenString, errAccess := accessToken.generateTokenString(user.Email)
	refreshTokenString, errRefresh := refreshToken.generateTokenString(user.Email)

	if errAccess != nil || errRefresh != nil {
		http.Error(w, "Error generating token string", http.StatusInternalServerError)
		return
	}

	accessToken.saveTokenAsCookie(w, accessTokenString)
	refreshToken.saveTokenAsCookie(w, refreshTokenString)

	response := map[string]string{"message": "register successful"}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}
