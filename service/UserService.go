package service

import (
	"bytes"
	"encoding/json"
	"github.com/AliSahib998/QuotesAssesments/errhandler"
	"github.com/AliSahib998/QuotesAssesments/model"
	"github.com/AliSahib998/QuotesAssesments/repo"
	"github.com/AliSahib998/QuotesAssesments/validator"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type IUserService interface {
	CreateUser(user model.UserRegistration) error
	Login(req model.Login) (*model.LoginResponse, error)
}

type UserService struct {
}

func (u *UserService) CreateUser(user model.UserRegistration) error {
	err := validator.Validation(user)
	if err != nil {
		return err
	}

	var query = model.SearchQuery{
		QueryString:  user.Username,
		SearchField:  "username",
		SortField:    "",
		SortOrder:    "",
		IsFullSearch: false,
	}

	users, err := repo.SearchUser(&query, "user")

	if err != nil {
		return err
	}

	if len(users) > 0 {
		return errhandler.NewBadRequestError("username was exist", nil)
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	var userDocument = model.UserDocument{
		Id:          uuid.New().String(),
		Username:    user.Username,
		Password:    string(hashedPassword),
		Name:        user.Name,
		Surname:     user.Surname,
		CreatedDate: time.Now(),
	}

	b := new(bytes.Buffer)
	err = json.NewEncoder(b).Encode(userDocument)
	if err != nil {
		return err
	}
	return repo.SaveDocument("user", b, userDocument.Id)
}

func (u *UserService) Login(req model.Login) (*model.LoginResponse, error) {
	var query = model.SearchQuery{
		QueryString:  req.Username,
		SearchField:  "username",
		SortField:    "",
		SortOrder:    "",
		IsFullSearch: false,
	}

	users, err := repo.SearchUser(&query, "user")
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(users[0].Password), []byte(req.Password))
	if err != nil {
		return nil, errhandler.NewBadRequestError("username or password was wrong", nil)
	}

	tk := model.Claim{Username: req.Username, Timestamp: time.Now()}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("TOKEN_KEY")))

	return &model.LoginResponse{AccessToken: tokenString}, nil
}
