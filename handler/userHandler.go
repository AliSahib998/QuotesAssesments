package handler

import (
	"encoding/json"
	"github.com/AliSahib998/QuotesAssesments/errhandler"
	"github.com/AliSahib998/QuotesAssesments/model"
	"github.com/AliSahib998/QuotesAssesments/service"
	"github.com/AliSahib998/QuotesAssesments/util"
	"github.com/go-chi/chi"
	"net/http"
)

type UserHandler struct {
	userService service.IUserService
}

func NewUserHandler(router *chi.Mux) *UserHandler {
	h := &UserHandler{
		userService: &service.UserService{},
	}
	router.Post("/user/register", errhandler.ErrorHandler(h.saveUser))
	router.Post("/user/login", errhandler.ErrorHandler(h.login))
	return h
}

func (u *UserHandler) saveUser(w http.ResponseWriter, r *http.Request) error {
	request := new(model.UserRegistration)
	err := util.ParseRequest(r, request)
	if err != nil {
		return err
	}
	err = u.userService.CreateUser(*request)
	if err != nil {
		return err
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	return nil
}

func (u *UserHandler) login(w http.ResponseWriter, r *http.Request) error {
	request := new(model.Login)
	err := util.ParseRequest(r, request)
	if err != nil {
		return err
	}
	resp, err := u.userService.Login(*request)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	return err
}
