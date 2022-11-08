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

func (u *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var req domain.UserUpdateParams

	err := ReadJSON(r.Body, &req)
	if err != nil {
		responses.ErrorBadRequestResponse(w, err.Error())
		return
	}

	currentUser, err := GetUser(r)
	if err != nil {
		responses.ErrorInternalServerResponse(w, err.Error())
		return
	}

	user, err := u.svc.UpdateUser(r.Context(), currentUser.ID, req)
	if err != nil {
		responses.ErrorInternalServerResponse(w, "Error updating user.")
		return
	}

	responses.SuccessResponse(w, user)
}

func (u *UserController) Check(w http.ResponseWriter, r *http.Request) {
	user, err := GetUser(r)
	if err != nil {
		responses.ErrorInternalServerResponse(w, err.Error())
		return
	}

	responses.SuccessResponse(w, user)
}

func (u *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	user, err := GetUser(r)
	if err != nil {
		responses.ErrorInternalServerResponse(w, err.Error())
	}

	err = u.svc.DeleteUser(r.Context(), user.ID)
	responses.SuccessResponse(w, nil)
}
