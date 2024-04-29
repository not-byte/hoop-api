package utils

import "net/http"

func HandleCookieError(w http.ResponseWriter, err error, tokenType string) {
	if err == http.ErrNoCookie {
		http.Error(w, "Missing "+tokenType, http.StatusUnauthorized)
	} else {
		http.Error(w, "Invalid request for "+tokenType, http.StatusBadRequest)
	}
}
