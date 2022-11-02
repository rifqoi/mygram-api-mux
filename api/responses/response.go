package responses

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
	Error   any    `json:"error,omitempty"`
}

func BadRequestResponse(w http.ResponseWriter, msg string, err any) {
	resp := Response{
		Message: msg,
		Error:   err,
	}
	baseResponse(w, resp, http.StatusBadRequest)
}

func SuccessResponse(w http.ResponseWriter, data any, err any) {
	resp := Response{
		Message: "SUCCESS",
		Data:    data,
		Error:   err,
	}
	baseResponse(w, resp, http.StatusOK)
}

func baseResponse(w http.ResponseWriter, res Response, status int) {
	w.Header().Set("Content-Type", "application/json")

	content, err := json.MarshalIndent(res, "", "    ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error parsing json"))
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(content); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error returning response"))
	}
}
