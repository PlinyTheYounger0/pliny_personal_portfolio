package handlers

import (
	"encoding/json"
	"net/http"
)

func (cfg *ApiConfig) respondWithError(w http.ResponseWriter, code int, msg string, err error) {
	type errorResp struct {
		Error string `json:"error"`
		Detail string `json:"detail,omitempty"`
	}

	resp := errorResp{
		Error: msg,
	}

	if err != nil {
		resp.Detail = err.Error()
	}

	dat, err := json.Marshal(msg)
	if err != nil {
		cfg.Logger.Error("Error Marshaling JSON", "error", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
