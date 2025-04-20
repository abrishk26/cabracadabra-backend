package utils

import (
	"encoding/json"
	"maps"
	"net/http"

	"github.com/google/uuid"
)

type envelop map[string]any

func CreateID() string {
	return uuid.New().String()
}

func WriteJSON(w http.ResponseWriter, status int, message envelop, headers http.Header) error {
	res, err := json.Marshal(message)
	if err != nil {
		return err
	}

	res = append(res, '\n')

	maps.Copy(w.Header(), headers)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(res)

	return nil
}

