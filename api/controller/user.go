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
		responses.ErrorBadRequestResponse(w, err.Error())
		return
	}

	errs := Validate(&req)
	if errs != nil {
		responses.ErrorBadRequestResponse(w, errs)
		return
	}

	res, err := u.svc.RegisterUser(r.Context(), req)
	if err != nil {
		responses.ErrorInternalServerResponse(w, err.Error())
		return
	}

	responses.SuccessResponse(w, res)
}

func (u *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var req domain.UserLoginParams

	err := ReadJSON(r.Body, &req)
	if err != nil {
		responses.ErrorBadRequestResponse(w, err.Error())
		return
	}

	errs := Validate(&req)
	if errs != nil {
		responses.ErrorBadRequestResponse(w, errs)
		return
	}

	token, err := u.svc.Login(r.Context(), req)
	if err != nil {
		responses.ErrorInternalServerResponse(w, err.Error())
		return
	}

	responses.SuccessResponse(w, token)
}

func (u *UserController) Check(w http.ResponseWriter, r *http.Request) {
	userInfo := r.Context().Value("userInfo")
	if userInfo == nil {
		responses.UnauthorizedRequest(w, "no user found")
		return
	}

	user, ok := userInfo.(*domain.User)
	if !ok {
		responses.ErrorInternalServerResponse(w, "casting error")
		return
	}

	responses.SuccessResponse(w, user)
}
