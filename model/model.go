package model

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Quote struct {
	Id           string   `json:"_id"`
	Content      string   `json:"content"`
	Tags         []string `json:"tags"`
	AuthorSlug   string   `json:"authorSlug"`
	Length       int      `json:"length"`
	DateAdded    string   `json:"dateAdded"`
	DateModified string   `json:"dateModified"`
}

type UserRegistration struct {
	Username string `json:"username"  validate:"required,gte=5,lte=15"`
	Password string `json:"password"  validate:"required,gte=7"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}

type SearchQuery struct {
	QueryString  string `json:"queryString"`
	SearchField  string `json:"searchField"`
	SortField    string `json:"sortField"`
	SortOrder    string `json:"sortOrder"`
	IsFullSearch bool   `json:"isFullSearch"`
}

type Claim struct {
	Username  string
	Timestamp time.Time
	jwt.StandardClaims
}

type QuoteDocument struct {
	Id         string `json:"id"`
	AuthorSlug string `json:"authorSlug"`
	Content    string `json:"content"`
	LikeCount  int    `json:"likeCount"`
}

type UserDocument struct {
	Id          string    `json:"id"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Name        string    `json:"name"`
	Surname     string    `json:"surname"`
	CreatedDate time.Time `json:"createdDate"`
}
