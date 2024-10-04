package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("Missing request body")

	}
	return json.NewDecoder(r.Body).Decode(payload)

}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)

}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})

}

func CeilHundred(x int) int {
	if x%100 != 0 {
		return x + 100 - (x % 100)
	}
	return x
}

func CeilTen(x int) int {
	if x%10 != 0 {
		return x + 10 - (x % 10)
	}
	return x
}
