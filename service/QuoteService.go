package service

import (
	"bytes"
	"encoding/json"
	"github.com/AliSahib998/QuotesAssesments/client"
	"github.com/AliSahib998/QuotesAssesments/model"
	"github.com/AliSahib998/QuotesAssesments/repo"
)

type IQuoteService interface {
	GetRandomQuote() (*model.QuoteDocument, error)
	LikeQuote(id string) error
	SearchQuote(query *model.SearchQuery) ([]model.QuoteDocument, error)
}

type QuoteService struct {
	QuoteClient client.IQuoteClient
}

func (q *QuoteService) GetRandomQuote() (*model.QuoteDocument, error) {
	resp, err := q.QuoteClient.GetRandomQuote()
	var quoteDocument = model.QuoteDocument{
		Id:         resp.Id,
		AuthorSlug: resp.AuthorSlug,
		Content:    resp.Content,
		LikeCount:  0,
	}

	b := new(bytes.Buffer)
	err = json.NewEncoder(b).Encode(quoteDocument)
	if err != nil {
		return nil, err
	}
	err = repo.Save("quote", b)

	if err != nil {
		return nil, err
	}

	return &quoteDocument, err
}

func (q *QuoteService) LikeQuote(id string) error {
	var query = model.SearchQuery{
		QueryString:      id,
		SearchField:      "id",
		SortField:        "",
		SortOrder:        "",
		IsWildCardSearch: false,
	}

	quotes, err := repo.SearchQuote(&query, "quote")

	if err != nil {
		return err
	}

	for _, v := range quotes {
		v.LikeCount = v.LikeCount + 1
		var buf bytes.Buffer
		jsonData, _ := json.Marshal(v)
		_, err = buf.Write(jsonData)
		_ = repo.Save("quote", &buf)
	}

	return nil
}

func (q *QuoteService) SearchQuote(query *model.SearchQuery) ([]model.QuoteDocument, error) {

	quotes, err := repo.SearchQuote(query, "quote")

	if err != nil {
		return nil, err
	}

	return quotes, nil
}
