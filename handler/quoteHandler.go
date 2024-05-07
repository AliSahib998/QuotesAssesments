package handler

import (
	"encoding/json"
	"fmt"
	"github.com/AliSahib998/QuotesAssesments/client"
	"github.com/AliSahib998/QuotesAssesments/errhandler"
	"github.com/AliSahib998/QuotesAssesments/model"
	"github.com/AliSahib998/QuotesAssesments/service"
	"github.com/AliSahib998/QuotesAssesments/util"
	"github.com/go-chi/chi"
	"github.com/go-resty/resty/v2"
	"net/http"
)

type QuoteHandler struct {
	quoteService service.IQuoteService
}

func NewQuoteHandler(router *chi.Mux) *QuoteHandler {
	h := &QuoteHandler{
		quoteService: &service.QuoteService{
			QuoteClient: &client.QuoteClient{
				Client: resty.New(),
			},
		},
	}
	router.Get("/quote", errhandler.ErrorHandler(h.GetQuote))
	router.Post("/quote/{id}", errhandler.ErrorHandler(h.LikeQuote))
	router.Post("/quote/search", errhandler.ErrorHandler(h.SearchQuote))
	return h
}

func (q *QuoteHandler) GetQuote(w http.ResponseWriter, r *http.Request) error {
	priority := r.URL.Query().Get("priority")
	resp, err := q.quoteService.GetQuote(priority)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	return err
}

func (q *QuoteHandler) LikeQuote(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")
	if len(id) == 0 {
		return errhandler.NewBadRequestError(fmt.Sprintf("%s", "invalid id"), nil)
	}
	err := q.quoteService.LikeQuote(id)
	w.Header().Set("Content-Type", "application/json")
	return err
}

func (q *QuoteHandler) SearchQuote(w http.ResponseWriter, r *http.Request) error {
	request := new(model.SearchQuery)
	err := util.ParseRequest(r, request)
	if err != nil {
		return err
	}
	resp, err := q.quoteService.SearchQuote(request)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	return err
}
