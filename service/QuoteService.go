package service

import (
	"bytes"
	"encoding/json"
	"github.com/AliSahib998/QuotesAssesments/client"
	"github.com/AliSahib998/QuotesAssesments/errhandler"
	"github.com/AliSahib998/QuotesAssesments/model"
	"github.com/AliSahib998/QuotesAssesments/repo"
)

type IQuoteService interface {
	GetQuote(priority string) (*model.QuoteDocument, error)
	GetQuoteById(id string) (*model.QuoteDocument, error)
	LikeQuote(id string) error
	SearchQuote(query *model.SearchQuery) ([]model.QuoteDocument, error)
}

type QuoteService struct {
	QuoteClient client.IQuoteClient
}

func (q *QuoteService) GetQuote(priority string) (*model.QuoteDocument, error) {
	switch priority {
	case "low", "high":
		return getQuoteFromStorageByLikeCount(priority)
	default:
		return getQuoteFromClient(q)
	}
}

func (q *QuoteService) LikeQuote(id string) error {
	var query = model.SearchQuery{
		QueryString:  id,
		SearchField:  "id",
		SortField:    "",
		SortOrder:    "",
		IsFullSearch: false,
	}

	quotes, err := repo.SearchQuote(&query, "quote")

	if err != nil {
		return err
	}

	if len(quotes) > 0 {
		var quote = quotes[0]
		quote.LikeCount = quote.LikeCount + 1
		var buf bytes.Buffer
		jsonData, err := json.Marshal(quote)
		if err != nil {
			return err
		}
		_, err = buf.Write(jsonData)
		if err != nil {
			return err
		}
		err = repo.UpdateDocument("quote", &buf, id)
	}

	return err
}

func (q *QuoteService) SearchQuote(query *model.SearchQuery) ([]model.QuoteDocument, error) {

	quotes, err := repo.SearchQuote(query, "quote")

	if err != nil {
		return nil, err
	}

	return quotes, nil
}

func (q *QuoteService) GetQuoteById(id string) (*model.QuoteDocument, error) {
	var query = model.SearchQuery{
		QueryString:  id,
		SearchField:  "id",
		SortField:    "",
		SortOrder:    "",
		IsFullSearch: false,
	}

	quotes, err := repo.SearchQuote(&query, "quote")

	if err != nil {
		return nil, err
	}

	if len(quotes) > 0 {
		var quote = quotes[0]
		return &quote, nil
	}

	return nil, errhandler.NewNotFoundError("the quote was not found", nil)
}

func getQuoteFromStorageByLikeCount(priority string) (*model.QuoteDocument, error) {
	var sortOrder = "asc"
	if priority == "high" {
		sortOrder = "desc"
	}
	var query = model.SearchQuery{
		QueryString:  "",
		SearchField:  "",
		SortField:    "likeCount",
		SortOrder:    sortOrder,
		IsFullSearch: false,
	}

	quotes, err := repo.SearchQuote(&query, "quote")

	if err != nil {
		return nil, err
	}

	if len(quotes) > 0 {
		return &quotes[0], nil
	}

	return nil, errhandler.NewNotFoundError("quote was not found in database with parameter", nil)
}

func getQuoteFromClient(q *QuoteService) (*model.QuoteDocument, error) {
	resp, err := q.QuoteClient.GetRandomQuote()
	var quoteDocument = model.QuoteDocument{
		Id:         resp.Id,
		AuthorSlug: resp.AuthorSlug,
		Content:    resp.Content,
		LikeCount:  0,
	}
	saveDocument(&quoteDocument)
	return &quoteDocument, err
}

func saveDocument(quoteDocument *model.QuoteDocument) {
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(quoteDocument)
	if err != nil {
		return
	}
	err = repo.SaveDocument("quote", b, quoteDocument.Id)

	if err != nil {
		return
	}
}
