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
	handler.NewHealthHandler(router)
	handler.NewQuoteHandler(router)
	handler.NewUserHandler(router)
	config.LoadESClient()
	port := strconv.Itoa(config.Props.Port)
	log.Info("Starting server at port: ", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

/*

  add scheduler to save quote in repo search
  config repo search with go


   1) add register api
   2) add login api
   3) call get quote endpoint -> accept query param (random, high_rate)
   4) expose like endpoint
   5) implement grpc and graphql
   6) implement search api with quotes (repo)

*/
