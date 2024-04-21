package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) error {
	response, err := json.Marshal(payload)

	if err != nil {
		respondErr := RespondWithError(w,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError))
		if respondErr != nil {
			return respondErr
		}

		return fmt.Errorf("failed to marshall payload: %w", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if _, err := w.Write(response); err != nil {
		return fmt.Errorf("failed to write response: %w", err)
	}
	return nil
}

func RespondWithError(w http.ResponseWriter, code int, message string) error {
	return RespondWithJSON(w,
		code,
		struct {
			Error string `json:"error"`
		}{
			Error: message,
		})
}

func SuccessResponseWith200(w http.ResponseWriter, payload interface{}) error {
	return RespondWithJSON(w,
		http.StatusOK,
		payload)
}

func SuccessResponseWith201(w http.ResponseWriter, payload interface{}) error {
	return RespondWithJSON(w,
		http.StatusCreated,
		payload)
}

func RespondWith400(w http.ResponseWriter, message string) error {
	return RespondWithError(w,
		http.StatusBadRequest,
		message)
}

func RespondWith401(w http.ResponseWriter, message string) error {
	return RespondWithError(w,
		http.StatusUnauthorized,
		message)
}

func RespondWith404(w http.ResponseWriter) error {
	return RespondWithError(w,
		http.StatusNotFound,
		http.StatusText(http.StatusNotFound))
}

func RespondWith500(w http.ResponseWriter) error {
	return RespondWithError(w,
		http.StatusInternalServerError,
		http.StatusText(http.StatusInternalServerError))
}
