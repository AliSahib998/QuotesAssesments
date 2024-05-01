package main

import (
	"github.com/AliSahib998/QuotesAssesments/config"
	"github.com/AliSahib998/QuotesAssesments/handler"
	"github.com/AliSahib998/QuotesAssesments/middleware"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func main() {
	config.LoadConfig()
	router := chi.NewRouter()
	router.Use(middleware.JwtAuthentication)
	router.Use(middleware.RequestParamsMiddleware)
	handler.NewQuoteHandler(router)
	handler.NewUserHandler(router)
	config.LoadESClient()
	port := strconv.Itoa(config.Props.Port)
	log.Info("Starting server at port: ", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
