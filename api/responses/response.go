package responses

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   any    `json:"error,omitempty"`
}

func ErrorBadRequestResponse(w http.ResponseWriter, err any) {
	resp := Response{
		Message: "BAD_REQUEST_ERROR",
		Error:   err,
	}
	baseResponse(w, resp, http.StatusBadRequest)
}

func ErrorInternalServerResponse(w http.ResponseWriter, err any) {
	resp := Response{
		Message: "INTERNAL_SERVER_ERROR",
		Error:   err,
	}
	baseResponse(w, resp, http.StatusInternalServerError)
}

func ErrorUnprocessableEntity(w http.ResponseWriter, err any) {
	resp := Response{
		Message: "UNPROCESSABLE_ENTITY",
		Error:   err,
	}
	baseResponse(w, resp, http.StatusUnprocessableEntity)
}

func UnauthorizedRequest(w http.ResponseWriter, err any) {
	resp := Response{
		Message: "UNAUTHORIZED_REQUEST",
		Error:   err,
	}

	baseResponse(w, resp, http.StatusUnauthorized)
}

func SuccessResponse(w http.ResponseWriter, data any) {
	resp := Response{
		Message: "SUCCESS",
		Data:    data,
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
