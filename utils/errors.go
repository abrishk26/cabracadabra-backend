package utils

import (
	"net/http"
)

func errorResponse(w http.ResponseWriter, _ *http.Request, status int, message any) {
	env := envelop{"error": message}

	err := WriteJSON(w, status, env, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func ServerErrorResponse(w http.ResponseWriter, r *http.Request) {
	message := "the server encountered a problem and could not process your request"

	errorResponse(w, r, http.StatusInternalServerError, message)
}