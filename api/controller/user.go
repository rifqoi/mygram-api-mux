package controller

import (
	"net/http"

	"github.com/rifqoi/mygram-api-mux/api/responses"
	"github.com/rifqoi/mygram-api-mux/domain"
	"github.com/rifqoi/mygram-api-mux/services"
)

type UserController struct {
	svc *services.UserService
}

func NewUserController(svc *services.UserService) *UserController {
	return &UserController{
		svc: svc,
	}
}

func (u *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req domain.UserCreateParams

	err := ReadJSON(r.Body, &req)
	if err != nil {
		responses.ErrorBadRequestResponse(w, "BAD_REQUEST", err.Error())
		return
	}

	res, err := u.svc.InsertUser(r.Context(), req)
	if err != nil {
		responses.ErrorInternalServerResponse(w, "INTERNAL_SERVER_ERROR", err.Error())
		return
	}

	responses.SuccessResponse(w, res)
}
