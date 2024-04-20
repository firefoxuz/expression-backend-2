package utils

import (
	"encoding/json"
	"net/http"
)

func DecodeBody(w http.ResponseWriter, r *http.Request, request interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return err
	}
	return nil
}
