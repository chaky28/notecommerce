package server

import (
	"errors"
	"net/http"
)

type ValidRequest struct {
	Method string
	Params []string
}

func ValidateRequest(w http.ResponseWriter, r *http.Request, vr ValidRequest) error {
	if err := validateMethod(w, r, vr.Method); err != nil {
		return err
	}
	if err := validateParams(w, r, vr.Params); err != nil {
		return err
	}
	return nil
}

func validateParams(w http.ResponseWriter, r *http.Request, params []string) error {
	for _, param := range params {
		if r.URL.Query().Get(param) == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("ERROR: " + param + " parameter missing and required in URL"))
			return errors.New("missing required parameter")
		}
	}
	return nil
}

func validateMethod(w http.ResponseWriter, r *http.Request, method string) error {
	if r.Method != method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("ERROR: Method not allowed"))
		return errors.New("method not allowed")
	}
	return nil
}
